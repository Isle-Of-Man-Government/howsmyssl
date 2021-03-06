// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/documentation.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Documentation` provides the information for describing a service.
//
// Example:
// <pre><code>documentation:
//   summary: >
//     The Google Calendar API gives access
//     to most calendar features.
//   pages:
//   - name: Overview
//     content: &#40;== include google/foo/overview.md ==&#41;
//   - name: Tutorial
//     content: &#40;== include google/foo/tutorial.md ==&#41;
//     subpages;
//     - name: Java
//       content: &#40;== include google/foo/tutorial_java.md ==&#41;
//   rules:
//   - selector: google.calendar.Calendar.Get
//     description: >
//       ...
//   - selector: google.calendar.Calendar.Put
//     description: >
//       ...
// </code></pre>
// Documentation is provided in markdown syntax. In addition to
// standard markdown features, definition lists, tables and fenced
// code blocks are supported. Section headers can be provided and are
// interpreted relative to the section nesting of the context where
// a documentation fragment is embedded.
//
// Documentation from the IDL is merged with documentation defined
// via the config at normalization time, where documentation provided
// by config rules overrides IDL provided.
//
// A number of constructs specific to the API platform are supported
// in documentation text.
//
// In order to reference a proto element, the following
// notation can be used:
// <pre><code>&#91;fully.qualified.proto.name]&#91;]</code></pre>
// To override the display text used for the link, this can be used:
// <pre><code>&#91;display text]&#91;fully.qualified.proto.name]</code></pre>
// Text can be excluded from doc using the following notation:
// <pre><code>&#40;-- internal comment --&#41;</code></pre>
// Comments can be made conditional using a visibility label. The below
// text will be only rendered if the `BETA` label is available:
// <pre><code>&#40;--BETA: comment for BETA users --&#41;</code></pre>
// A few directives are available in documentation. Note that
// directives must appear on a single line to be properly
// identified. The `include` directive includes a markdown file from
// an external source:
// <pre><code>&#40;== include path/to/file ==&#41;</code></pre>
// The `resource_for` directive marks a message to be the resource of
// a collection in REST view. If it is not specified, tools attempt
// to infer the resource from the operations in a collection:
// <pre><code>&#40;== resource_for v1.shelves.books ==&#41;</code></pre>
// The directive `suppress_warning` does not directly affect documentation
// and is documented together with service config validation.
type Documentation struct {
	// A short summary of what the service does. Can only be provided by
	// plain text.
	Summary string `protobuf:"bytes,1,opt,name=summary" json:"summary,omitempty"`
	// The top level pages for the documentation set.
	Pages []*Page `protobuf:"bytes,5,rep,name=pages" json:"pages,omitempty"`
	// Documentation rules for individual elements of the service.
	Rules []*DocumentationRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// The URL to the root of documentation.
	DocumentationRootUrl string `protobuf:"bytes,4,opt,name=documentation_root_url,json=documentationRootUrl" json:"documentation_root_url,omitempty"`
	// Declares a single overview page. For example:
	// <pre><code>documentation:
	//   summary: ...
	//   overview: &#40;== include overview.md ==&#41;
	// </code></pre>
	// This is a shortcut for the following declaration (using pages style):
	// <pre><code>documentation:
	//   summary: ...
	//   pages:
	//   - name: Overview
	//     content: &#40;== include overview.md ==&#41;
	// </code></pre>
	// Note: you cannot specify both `overview` field and `pages` field.
	Overview string `protobuf:"bytes,2,opt,name=overview" json:"overview,omitempty"`
}

func (m *Documentation) Reset()                    { *m = Documentation{} }
func (m *Documentation) String() string            { return proto.CompactTextString(m) }
func (*Documentation) ProtoMessage()               {}
func (*Documentation) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Documentation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *Documentation) GetRules() []*DocumentationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A documentation rule provides information about individual API elements.
type DocumentationRule struct {
	// The selector is a comma-separated list of patterns. Each pattern is a
	// qualified name of the element which may end in "*", indicating a wildcard.
	// Wildcards are only allowed at the end and for a whole component of the
	// qualified name, i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". To
	// specify a default for all applicable elements, the whole pattern "*"
	// is used.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Description of the selected API(s).
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Deprecation description of the selected element(s). It can be provided if an
	// element is marked as `deprecated`.
	DeprecationDescription string `protobuf:"bytes,3,opt,name=deprecation_description,json=deprecationDescription" json:"deprecation_description,omitempty"`
}

func (m *DocumentationRule) Reset()                    { *m = DocumentationRule{} }
func (m *DocumentationRule) String() string            { return proto.CompactTextString(m) }
func (*DocumentationRule) ProtoMessage()               {}
func (*DocumentationRule) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

// Represents a documentation page. A page can contain subpages to represent
// nested documentation set structure.
type Page struct {
	// The name of the page. It will be used as an identity of the page to
	// generate URI of the page, text of the link to this page in navigation,
	// etc. The full page name (start from the root page name to this page
	// concatenated with `.`) can be used as reference to the page in your
	// documentation. For example:
	// <pre><code>pages:
	// - name: Tutorial
	//   content: &#40;== include tutorial.md ==&#41;
	//   subpages:
	//   - name: Java
	//     content: &#40;== include tutorial_java.md ==&#41;
	// </code></pre>
	// You can reference `Java` page using Markdown reference link syntax:
	// `[Java][Tutorial.Java]`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The Markdown content of the page. You can use <code>&#40;== include {path} ==&#41;</code>
	// to include content from a Markdown file.
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	// Subpages of this page. The order of subpages specified here will be
	// honored in the generated docset.
	Subpages []*Page `protobuf:"bytes,3,rep,name=subpages" json:"subpages,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Page) GetSubpages() []*Page {
	if m != nil {
		return m.Subpages
	}
	return nil
}

func init() {
	proto.RegisterType((*Documentation)(nil), "google.api.Documentation")
	proto.RegisterType((*DocumentationRule)(nil), "google.api.DocumentationRule")
	proto.RegisterType((*Page)(nil), "google.api.Page")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/documentation.proto", fileDescriptor7)
}

var fileDescriptor7 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0x3f, 0xed, 0x5f, 0x9d, 0xa2, 0xe8, 0x20, 0x35, 0x08, 0x42, 0xe9, 0xa1, 0xf4,
	0x20, 0x09, 0x58, 0xc1, 0x7b, 0xe9, 0xcd, 0x4b, 0x09, 0x78, 0x2e, 0xe9, 0x76, 0x5c, 0x02, 0x49,
	0x26, 0x6c, 0x92, 0x8a, 0xaf, 0xe0, 0x13, 0xfa, 0x38, 0xee, 0x6e, 0xd2, 0x76, 0x83, 0x78, 0x09,
	0x99, 0x7c, 0xbf, 0x9d, 0xf9, 0xe6, 0xdb, 0xc0, 0xab, 0x64, 0x96, 0x19, 0x85, 0x92, 0xb3, 0xa4,
	0x90, 0x21, 0x2b, 0x19, 0x49, 0x2a, 0x4a, 0xc5, 0x35, 0x47, 0xad, 0x94, 0x94, 0x69, 0x15, 0xe9,
	0x47, 0x54, 0x91, 0xda, 0xa7, 0x82, 0x04, 0x17, 0xef, 0xa9, 0x8c, 0x76, 0x2c, 0x9a, 0x9c, 0x8a,
	0x3a, 0xa9, 0x53, 0x2e, 0x42, 0x7b, 0x00, 0xa1, 0x6b, 0xa6, 0xe9, 0xe9, 0xb7, 0x07, 0x97, 0x2b,
	0x97, 0xc1, 0x00, 0xce, 0xaa, 0x26, 0xcf, 0x13, 0xf5, 0x19, 0x78, 0x13, 0x6f, 0x7e, 0x11, 0x1f,
	0x4a, 0x9c, 0xc1, 0xb0, 0x4c, 0x24, 0x55, 0xc1, 0x70, 0xe2, 0xcf, 0x47, 0x4f, 0xd7, 0xe1, 0xa9,
	0x4f, 0xb8, 0xd6, 0x42, 0xdc, 0xca, 0xb8, 0x80, 0xa1, 0x6a, 0x32, 0xcd, 0xf9, 0x96, 0x7b, 0x70,
	0xb9, 0xde, 0xac, 0x58, 0x53, 0x71, 0xcb, 0xe2, 0x33, 0x8c, 0x7b, 0x5e, 0x37, 0x8a, 0xb9, 0xde,
	0x34, 0x2a, 0x0b, 0x06, 0xd6, 0xc5, 0x6d, 0x4f, 0x8d, 0xb5, 0xf8, 0xa6, 0x32, 0xbc, 0x87, 0x73,
	0xde, 0x9b, 0x85, 0xe9, 0x23, 0xf8, 0x67, 0xb9, 0x63, 0x3d, 0xfd, 0xf2, 0xe0, 0xe6, 0xd7, 0x38,
	0x73, 0xa2, 0xa2, 0x8c, 0x44, 0xcd, 0xaa, 0xdb, 0xef, 0x58, 0xe3, 0x04, 0x46, 0x3b, 0xaa, 0x84,
	0x4a, 0x4b, 0x83, 0x77, 0x0d, 0xdd, 0x4f, 0xf8, 0x02, 0x77, 0x3b, 0x2a, 0x15, 0x89, 0xd6, 0xa3,
	0x4b, 0xfb, 0x96, 0x1e, 0x3b, 0xf2, 0xea, 0xa4, 0x4e, 0xb7, 0x30, 0x30, 0x11, 0x21, 0xc2, 0xa0,
	0x48, 0x72, 0xea, 0x46, 0xdb, 0x77, 0x93, 0xb8, 0xbe, 0xad, 0x5a, 0xdb, 0xec, 0x46, 0x1e, 0x4a,
	0x7c, 0xd4, 0x66, 0x9b, 0x6d, 0x1b, 0xba, 0xff, 0x47, 0xe8, 0x47, 0x62, 0x39, 0x83, 0x2b, 0xc1,
	0xb9, 0x03, 0x2c, 0xb1, 0xb7, 0xff, 0xda, 0xdc, 0xfe, 0xda, 0xdb, 0xfe, 0xb7, 0xbf, 0xc1, 0xe2,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x33, 0x07, 0x1b, 0x82, 0x55, 0x02, 0x00, 0x00,
}
