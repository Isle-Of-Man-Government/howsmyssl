package main

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/logging"

	topk "github.com/dgryski/go-topk"
	"github.com/jmhodges/howsmyssl/domains"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type originAllower struct {
	ns       *expvar.Map
	hostname string
	gclog    logClient

	mu                 *sync.RWMutex
	blocked            map[string]bool
	topKAllDomains     *topk.Stream
	topKOfflistDomains *topk.Stream
}

type logClient interface {
	Log(logging.Entry)
	Flush()
}

func newOriginAllower(hostname string, gclog logClient, ns *expvar.Map) *originAllower {
	mu := &sync.RWMutex{}
	topKAllDomains := topk.New(100)
	topKOfflistDomains := topk.New(100)
	lifetime := new(expvar.Map).Init()
	ns.Set("lifetime", lifetime)
	lifetime.Set("top_all_domains", expvar.Func(func() interface{} {
		mu.RLock()
		defer mu.RUnlock()
		return topKAllDomains.Keys()
	}))
	lifetime.Set("top_offlist_domains", expvar.Func(func() interface{} {
		mu.RLock()
		defer mu.RUnlock()
		return topKOfflistDomains.Keys()
	}))

	oa := &originAllower{
		ns:                 ns,
		hostname:           hostname,
		gclog:              gclog,
		mu:                 mu,
		blocked:            make(map[string]bool),
		topKAllDomains:     topKAllDomains,
		topKOfflistDomains: topKOfflistDomains,
	}
	return oa
}

func (oa *originAllower) Allow(r *http.Request) (string, bool) {
	origin := r.Header.Get("Origin")
	referrer := r.Header.Get("Referer")

	apiKey := r.FormValue("key")
	userAgent := r.Header.Get("User-Agent")

	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("error splitting %#v as host:port: %s", r.RemoteAddr, err)
		remoteIP = "0.0.0.0"
	}
	entry := &apiLogEntry{
		DetectedDomain: "",
		Allowed:        false,
		APIKey:         apiKey,
		Headers: headers{
			Origin:    origin,
			Referrer:  referrer,
			UserAgent: userAgent,
		},
	}
	defer func() {
		go oa.countRequest(entry, r, remoteIP)
	}()

	domain, ok := oa.checkAllowedOriginAndReferrer(origin, referrer)
	entry.Allowed = ok
	entry.DetectedDomain = domain
	if !ok {
		entry.RejectionReason = rejectionConfig
	}
	return domain, ok
}

func (oa *originAllower) checkAllowedOriginAndReferrer(origin, referrer string) (string, bool) {
	if origin == "" && referrer == "" {
		return "", true
	}
	if origin != "" {
		return oa.isDomainAllowed(origin)
	}
	return oa.isDomainAllowed(referrer)
}

// isDomainAllowed checks if the detected domain from the request headers and
// whether domain is allowed to make requests against howsmyssl's API.
func (oa *originAllower) isDomainAllowed(d string) (string, bool) {
	domain, err := effectiveDomain(d)
	if err != nil {
		return "", false
	}
	isBlocked := oa.blocked[domain]
	return domain, !isBlocked
}

func (oa *originAllower) countRequest(entry *apiLogEntry, r *http.Request, remoteIP string) {
	oa.gclog.Log(logging.Entry{
		Payload:     entry,
		HTTPRequest: &logging.HTTPRequest{Request: r, RemoteIP: remoteIP},
		Labels: map[string]string{
			"server_hostname": oa.hostname,
			"app":             "howsmyssl",
		},
	})

	if entry.DetectedDomain == "" {
		return
	}

	oa.mu.Lock()
	defer oa.mu.Unlock()
	oa.topKAllDomains.Insert(entry.DetectedDomain, 1)
	if !entry.Allowed {
		oa.topKOfflistDomains.Insert(entry.DetectedDomain, 1)
	}
}

func effectiveDomain(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	host := u.Host
	if host == "" {
		return "", fmt.Errorf("unparsable domain string %#v", str)
	}
	i := strings.Index(host, ":")
	if i >= 0 {
		host = host[:i]
	}

	if host == "localhost" {
		return "localhost", nil
	}
	d, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return "", err
	}
	return d, nil
}

type originsConfig struct {
	// BlockedOrigins are domains that are not to be allowed as referrers to the
	// API. They should not have a scheme or path, but only the domain, as in
	// "example.com".
	BlockedOrigins []string `json:"blocked_origins"`
}

type rejectionReason string

const rejectionConfig = rejectionReason("config")

type apiLogEntry struct {
	DetectedDomain  string          `json:"detected_domain"`
	Allowed         bool            `json:"allowed"`
	APIKey          string          `json:"api_key"`
	RejectionReason rejectionReason `json:"rejection_reason"`
	Headers         headers         `json:"headers"`
}

type headers struct {
	Origin    string `json:"origin"`
	Referrer  string `json:"referrer"`
	UserAgent string `json:"user_agent"`
}

func loadGoogleServiceAccount(fp string) *googleConfig {
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("unable to read Google service account config %#v: %s", fp, err)
	}
	c := &googleConfig{}
	err = json.Unmarshal(bs, c)
	if err != nil {
		log.Fatalf("unable to parse project ID from Google service account config %#v: %s", fp, err)
	}
	if c.ProjectID == "" {
		log.Fatalf("blank project ID in Google service account config %#v: %s", fp, err)
	}
	jwtConf, err := google.JWTConfigFromJSON(bs, logging.WriteScope)
	if err != nil {
		log.Fatalf("unable to parse Google service account config %#v: %s", fp, err)
	}
	c.conf = jwtConf
	return c
}

type googleConfig struct {
	ProjectID string `json:"project_id"`

	conf *jwt.Config `json:"-"`
}

var _ logClient = nullLogClient{}

type nullLogClient struct{}

func (n nullLogClient) Log(e logging.Entry) {
}

func (n nullLogClient) Flush() {
}

func kickOffFetchBlockedDomainsForever(oa *originAllower, client domains.DomainCheckClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	err := fetchBlockedDomains(ctx, oa, client)
	if err != nil {
		log.Fatalf("unable to fetch allowed domains on boot: %s", err)
	}
	cancel()

	// We fork this goroutine here instead of just running all of the code in
	// this func in it so that the first fetchBlockedDomains set ups the
	// originAllower before traffic is served.
	go func() {
		tick := time.NewTicker(60 * time.Second)
		for range tick.C {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			err := fetchBlockedDomains(ctx, oa, client)
			cancel()
			if err != nil {
				log.Printf("unable to fetch allowed domains: %s", err)
			}
		}
	}()
}

func fetchBlockedDomains(ctx context.Context, oa *originAllower, client domains.DomainCheckClient) error {
	res, err := client.AllBlockedDomains(ctx, &domains.AllBlockedDomainsRequest{})
	if err != nil {
		return err
	}
	m := make(map[string]bool)
	for _, d := range res.Domains {
		m[d] = true
	}
	oa.mu.Lock()
	defer oa.mu.Unlock()
	oa.blocked = m
	return nil
}
