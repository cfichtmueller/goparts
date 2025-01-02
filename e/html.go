// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

import "strconv"

const (
	Stylesheet = "stylesheet"

	AttrAccept          = "accept"
	AttrAcceptCharset   = "accept-charset"
	AttrAccesskey       = "accesskey"
	AttrAction          = "action"
	AttrAlign           = "align"
	AttrAllow           = "allow"
	AttrAlt             = "alt"
	AttrAs              = "as"
	AttrAsync           = "async"
	AttrAutocapitalize  = "autocapitalize"
	AttrAutocomplete    = "autocomplete"
	AttrAutoplay        = "autoplay"
	AttrCapture         = "capture"
	AttrCharset         = "charset"
	AttrChecked         = "checked"
	AttrCite            = "cite"
	AttrClass           = "class"
	AttrCols            = "cols"
	AttrColspan         = "colspan"
	AttrContent         = "content"
	AttrContenteditable = "contenteditable"
	AttrControls        = "controls"
	AttrCoords          = "coords"
	AttrCrossorigin     = "crossorigin"
	AttrData            = "data"
	AttrDatatestid      = "data-testid"
	AttrDatetime        = "datetime"
	AttrDecoding        = "decoding"
	AttrDefault         = "default"
	AttrDefer           = "defer"
	AttrDir             = "dir"
	AttrDirname         = "dirname"
	AttrDisabled        = "disabled"
	AttrDownload        = "download"
	AttrDraggable       = "draggable"
	AttrEnctype         = "enctype"
	AttrEnterkeyhint    = "enterkeyhint"
	AttrFor             = "for"
	AttrForm            = "form"
	AttrFormaction      = "formaction"
	AttrFormenctype     = "formenctype"
	AttrFormmethod      = "formmethod"
	AttrFormnovalidate  = "formnovalidate"
	AttrFormtarget      = "formtarget"
	AttrHeaders         = "headers"
	AttrHeight          = "height"
	AttrHidden          = "hidden"
	AttrHigh            = "high"
	AttrHref            = "href"
	AttrHreflang        = "hreflang"
	AttrHttpEquiv       = "http-equiv"
	AttrId              = "id"
	AttrIntegrity       = "integrity"
	AttrIntrinsicsize   = "intrinsicsize"
	AttrInputmode       = "inputmode"
	AttrIsmap           = "ismap"
	AttrItemprop        = "itemprop"
	AttrKind            = "kind"
	AttrLabel           = "label"
	AttrLang            = "lang"
	AttrLanguage        = "language"
	AttrLoading         = "loading"
	AttrList            = "list"
	AttrLoop            = "loop"
	AttrLow             = "low"
	AttrMax             = "max"
	AttrMaxlength       = "maxlength"
	AttrMinlength       = "minlength"
	AttrMedia           = "media"
	AttrMethod          = "method"
	AttrMin             = "min"
	AttrMultiple        = "multiple"
	AttrMuted           = "muted"
	AttrName            = "name"
	AttrNovalidate      = "novalidate"
	AttrOpen            = "open"
	AttrOptimum         = "optimum"
	AttrPattern         = "pattern"
	AttrPing            = "ping"
	AttrPlaceholder     = "placeholder"
	AttrPlaysinline     = "playsinline"
	AttrPoster          = "poster"
	AttrPreload         = "preload"
	AttrReadonly        = "readonly"
	AttrReferrerpolicy  = "referrerpolicy"
	AttrRel             = "rel"
	AttrRequired        = "required"
	AttrReversed        = "reversed"
	AttrRole            = "role"
	AttrRows            = "rows"
	AttrRowspan         = "rowspan"
	AttrSandbox         = "sandbox"
	AttrScrope          = "scope"
	AttrScoped          = "scoped"
	AttrSelected        = "selected"
	AttrShape           = "shape"
	AttrSize            = "size"
	AttrSizes           = "sizes"
	AttrSlot            = "slot"
	AttrSpan            = "span"
	AttrSpellcheck      = "spellcheck"
	AttrSrc             = "src"
	AttrSrcdoc          = "srcdoc"
	AttrSrclang         = "srclang"
	Attrsrcset          = "srcset"
	Attrstart           = "start"
	AttrStep            = "step"
	AttrStyle           = "style"
	AttrSummary         = "sumary"
	AttrTabIndex        = "tabindex"
	AttrTarget          = "target"
	AttrTitle           = "title"
	AttrTranslate       = "translate"
	AttrType            = "type"
	AttrUsemap          = "usemap"
	AttrValue           = "value"
	AttrWidth           = "width"
	AttrWrap            = "wrap"

	AttrProperty = "property"
	AttrVersion  = "version"
	AttrXmlns    = "xmlns"

	AutoCompleteNewPassword = "new-password"
	AutoCompleteUsername    = "username"

	// Main root
	TagHtml = "html"

	// Document metadata
	TagBase  = "base"
	TagHead  = "head"
	TagLink  = "link"
	TagMeta  = "meta"
	TagStyle = "style"
	TagTitle = "title"

	// Sectioning root
	TagBody = "body"

	// Content sectioning
	TagAddress = "address"
	TagArticle = "article"
	TagAside   = "aside"
	TagFooter  = "footer"
	TagHeader  = "header"
	TagH1      = "h1"
	TagH2      = "h2"
	TagH3      = "h3"
	TagH4      = "h4"
	TagH5      = "h5"
	TagH6      = "h6"
	TagHgroup  = "hgroup"
	TagMain    = "main"
	TagNav     = "nav"
	TagSection = "section"
	TagSearch  = "search"

	// Text content
	TagBlockquote = "blockquote"
	TagDd         = "dd"
	TagDiv        = "div"
	TagDl         = "dl"
	TagDt         = "dt"
	TagFigcaption = "figcaption"
	TagFigure     = "figure"
	TagHr         = "hr"
	TagLi         = "li"
	TagMenu       = "menu"
	TagOl         = "ol"
	TagP          = "p"
	TagPre        = "pre"
	TagUl         = "ul"

	// Inline text semantics
	TagA      = "a"
	TagAbbr   = "abbr"
	TagB      = "b"
	TagBdi    = "bdi"
	TagBdo    = "bdo"
	TagBr     = "br"
	TagCite   = "cite"
	TagCode   = "code"
	TagData   = "data"
	TagDfn    = "dfn"
	TagEm     = "em"
	TagI      = "i"
	TagKbd    = "kbd"
	TagMark   = "mark"
	TagQ      = "q"
	TagRb     = "rb"
	TagRp     = "rp"
	TagRt     = "rt"
	TagRuby   = "ruby"
	TagS      = "s"
	TagSamp   = "samp"
	TagSmall  = "small"
	TagSpan   = "span"
	TagStrong = "strong"
	TagSub    = "sub"
	TagSup    = "sup"
	TagTime   = "time"
	TagU      = "u"
	TagVar    = "var"
	TagWbr    = "wbr"

	// Image and multimedia
	TagArea  = "area"
	TagAudio = "audio"
	TagImg   = "img"
	TagMap   = "map"
	TagTrack = "track"
	TagVideo = "video"

	// Embedded content
	TagEmbed       = "embed"
	TagFencedFrame = "fencedframe"
	TagIframe      = "iframe"
	TagObject      = "object"
	TagPicture     = "picture"
	TagPortal      = "portal"
	TagSource      = "source"

	// SVG and MathML
	TagSvg  = "svg"
	TagMath = "math"

	// Scripting
	TagCanvas   = "canvas"
	TagNoscript = "noscript"
	TagScript   = "script"

	// Demarcating edits
	TagDel = "del"
	TagIns = "ins"

	// Table content
	TagCaption  = "caption"
	TagCol      = "col"
	TagColgroup = "colgroup"
	TagTable    = "table"
	TagTbody    = "tbody"
	TagTd       = "td"
	TagTfoot    = "tfoot"
	TagTh       = "th"
	TagThead    = "thead"
	TagTr       = "tr"

	// Forms
	TagButton   = "button"
	TagDatalist = "datalist"
	TagFieldset = "fieldset"
	TagForm     = "form"
	TagInput    = "input"
	TagLabel    = "label"
	TagLegend   = "legend"
	TagMeter    = "meter"
	TagOptgroup = "optgroup"
	TagOption   = "option"
	TagOutput   = "output"
	TagProgress = "progress"
	TagSelect   = "select"
	TagTextarea = "textarea"

	// Interactive elements
	TagDetails = "details"
	TagDialog  = "dialog"
	TagSummary = "summary"

	// Web components
	TagSlot     = "slot"
	TagTemplate = "template"
)

var doctypeNode = Raw("<!DOCTYPE html>")

//
// Elements
//

// Main root

// Doctype creates a "<!DOCTYPE html>" node
func Doctype() Node {
	return doctypeNode
}

// Html creates an "html" node
func Html(children ...Node) Node {
	return &ElementNode{Tag: TagHtml, Children: children}
}

// Document metadata

// Base creates a "base" node
func Base(children ...Node) Node {
	return &ElementNode{Tag: TagBase, Children: children, AllowOmission: true}
}

// Head creates a "head" node
func Head(children ...Node) Node {
	return &ElementNode{Tag: TagHead, Children: children}
}

// Link creates a "link" node
func Link(children ...Node) Node {
	return &ElementNode{Tag: TagLink, Children: children, AllowOmission: true}
}

// Meta creates a "meta" node
func Meta(children ...Node) Node {
	return &ElementNode{Tag: TagMeta, Children: children, AllowOmission: true}
}

// Style creates a "style" node
func Style(children ...Node) Node {
	return &ElementNode{Tag: TagStyle, Children: children}
}

// Title creates a "title" node
func Title(children ...Node) Node {
	return &ElementNode{Tag: TagTitle, Children: children}
}

// Sectioning root

// Body creates a "body" node
func Body(children ...Node) Node {
	return &ElementNode{Tag: TagBody, Children: children}
}

// Content sectioning

// Address creates an "address" node
func Address(children ...Node) Node {
	return &ElementNode{Tag: TagAddress, Children: children}
}

// Article creates an "article" node
func Article(children ...Node) Node {
	return &ElementNode{Tag: TagArticle, Children: children}
}

// Aside creates an "aside" node
func Aside(children ...Node) Node {
	return &ElementNode{Tag: TagAside, Children: children}
}

// Footer creates a "footer" node
func Footer(children ...Node) Node {
	return &ElementNode{Tag: TagFooter, Children: children}
}

// Header creates a "header" node
func Header(children ...Node) Node {
	return &ElementNode{Tag: TagHeader, Children: children}
}

// H1 creates an "h1" node
func H1(children ...Node) Node {
	return &ElementNode{Tag: TagH1, Children: children}
}

// H2 creates an "h2" node
func H2(children ...Node) Node {
	return &ElementNode{Tag: TagH2, Children: children}
}

// H3 creates an "h3" node
func H3(children ...Node) Node {
	return &ElementNode{Tag: TagH3, Children: children}
}

// H4 creates an "h4" node
func H4(children ...Node) Node {
	return &ElementNode{Tag: TagH4, Children: children}
}

// H5 creates an "h5" node
func H5(children ...Node) Node {
	return &ElementNode{Tag: TagH5, Children: children}
}

// H6 creates an "h6" node
func H6(children ...Node) Node {
	return &ElementNode{Tag: TagH6, Children: children}
}

// HGroup creates an "hgroup" node
func Hgroup(children ...Node) Node {
	return &ElementNode{Tag: TagHgroup, Children: children}
}

// Main creates a "main" node
func Main(children ...Node) Node {
	return &ElementNode{Tag: TagMain, Children: children}
}

// Nav creates a "nav" node
func Nav(children ...Node) Node {
	return &ElementNode{Tag: TagNav, Children: children}
}

// Section creates a "section" node
func Section(children ...Node) Node {
	return &ElementNode{Tag: TagSection, Children: children}
}

// Search creates a "search" node
func Search(children ...Node) Node {
	return &ElementNode{Tag: TagSearch, Children: children}
}

// Text content

// Blockquote creates a "blockquote" node
func Blockquote(children ...Node) Node {
	return &ElementNode{Tag: TagBlockquote, Children: children}
}

// Dd creates a "dd" node
func Dd(children ...Node) Node {
	return &ElementNode{Tag: TagDd, Children: children}
}

// Div creates a "div" node
func Div(children ...Node) Node {
	return &ElementNode{Tag: TagDiv, Children: children}
}

// Dl creates a "dl" node
func Dl(children ...Node) Node {
	return &ElementNode{Tag: TagDl, Children: children}
}

// Dt creates a "dt" node
func Dt(children ...Node) Node {
	return &ElementNode{Tag: TagDt, Children: children}
}

// Figcaption creates a "figcaption" node
func Figcaption(children ...Node) Node {
	return &ElementNode{Tag: TagFigcaption, Children: children}
}

// Figure creates a "figure" node
func Figure(children ...Node) Node {
	return &ElementNode{Tag: TagFigure, Children: children}
}

// Hr creates a "hr" node
func Hr(children ...Node) Node {
	return &ElementNode{Tag: TagHr, Children: children, AllowOmission: true}
}

// Li creates a "li" node
func Li(children ...Node) Node {
	return &ElementNode{Tag: TagLi, Children: children}
}

// Menu creates a "menu" node
func Menu(children ...Node) Node {
	return &ElementNode{Tag: TagMenu, Children: children}
}

// Ol creates a "ol" node
func Ol(children ...Node) Node {
	return &ElementNode{Tag: TagOl, Children: children}
}

// P creates a "p" node
func P(children ...Node) Node {
	return &ElementNode{Tag: TagP, Children: children}
}

// Pre creates a "pre" node
func Pre(children ...Node) Node {
	return &ElementNode{Tag: TagPre, Children: children}
}

// Ul creates a "ul" node
func Ul(children ...Node) Node {
	return &ElementNode{Tag: TagUl, Children: children}
}

// Inline text semantics

// A creates an "a" node
func A(children ...Node) Node {
	return &ElementNode{Tag: TagA, Children: children}
}

// Abbr creates an "abbr" node
func Abbr(children ...Node) Node {
	return &ElementNode{Tag: TagAbbr, Children: children}
}

// B creates a "b" node
func B(children ...Node) Node {
	return &ElementNode{Tag: TagB, Children: children}
}

// Bdi creates a "bdi" node
func Bdi(children ...Node) Node {
	return &ElementNode{Tag: TagBdi, Children: children}
}

// Bdo creates a "bdo" node
func Bdo(children ...Node) Node {
	return &ElementNode{Tag: TagBdo, Children: children}
}

// Br creates a "br" node
func Br(children ...Node) Node {
	return &ElementNode{Tag: TagBr, Children: children, AllowOmission: true}
}

// Cite creates a "cite" node
func Cite(children ...Node) Node {
	return &ElementNode{Tag: TagCite, Children: children}
}

// Code creates a "code" node
func Code(children ...Node) Node {
	return &ElementNode{Tag: TagCode, Children: children}
}

// Data creates a "data" node
func Data(children ...Node) Node {
	return &ElementNode{Tag: TagData, Children: children}
}

// Dfn creates a "dfn" node
func Dfn(children ...Node) Node {
	return &ElementNode{Tag: TagDfn, Children: children}
}

// Em creates a "em" node
func Em(children ...Node) Node {
	return &ElementNode{Tag: TagEm, Children: children}
}

// I creates a "i" node
func I(children ...Node) Node {
	return &ElementNode{Tag: TagI, Children: children}
}

// Kbd creates a "kbd" node
func Kbd(children ...Node) Node {
	return &ElementNode{Tag: TagKbd, Children: children}
}

// Mark creates a "mark" node
func Mark(children ...Node) Node {
	return &ElementNode{Tag: TagMark, Children: children}
}

// Q creates a "q" node
func Q(children ...Node) Node {
	return &ElementNode{Tag: TagQ, Children: children}
}

// Rp creates a "rp" node
func Rp(children ...Node) Node {
	return &ElementNode{Tag: TagRp, Children: children}
}

// Rt creates a "rt" node
func Rt(children ...Node) Node {
	return &ElementNode{Tag: TagRt, Children: children}
}

// Ruby creates a "ruby" node
func Ruby(children ...Node) Node {
	return &ElementNode{Tag: TagRuby, Children: children}
}

// S creates a "s" node
func S(children ...Node) Node {
	return &ElementNode{Tag: TagS, Children: children}
}

// Samp creates a "samp" node
func Samp(children ...Node) Node {
	return &ElementNode{Tag: TagSamp, Children: children}
}

// Small creates a "small" node
func Small(children ...Node) Node {
	return &ElementNode{Tag: TagSmall, Children: children}
}

// Span creates a "span" node
func Span(children ...Node) Node {
	return &ElementNode{Tag: TagSpan, Children: children}
}

// Strong creates a "strong" node
func Strong(children ...Node) Node {
	return &ElementNode{Tag: TagStrong, Children: children}
}

// Sub creates a "sub" node
func Sub(children ...Node) Node {
	return &ElementNode{Tag: TagSub, Children: children}
}

// Sup creates a "sup" node
func Sup(children ...Node) Node {
	return &ElementNode{Tag: TagSup, Children: children}
}

// Time creates a "time" node
func Time(children ...Node) Node {
	return &ElementNode{Tag: TagTime, Children: children}
}

// U creates a "u" node
func U(children ...Node) Node {
	return &ElementNode{Tag: TagU, Children: children}
}

// Var creates a "var" node
func Var(children ...Node) Node {
	return &ElementNode{Tag: TagVar, Children: children}
}

// Wbr creates a "wbr" node
func Wbr(children ...Node) Node {
	return &ElementNode{Tag: TagWbr, Children: children, AllowOmission: true}
}

// Image and multimedia

// Area creates a "area" node
func Area(children ...Node) Node {
	return &ElementNode{Tag: TagArea, Children: children}
}

// Audio creates a "audio" node
func Audi(children ...Node) Node {
	return &ElementNode{Tag: TagAudio, Children: children}
}

// Img creates a "img" node
func Img(children ...Node) Node {
	return &ElementNode{Tag: TagImg, Children: children, AllowOmission: true}
}

// Map creates a "map" node
func Map(children ...Node) Node {
	return &ElementNode{Tag: TagMap, Children: children}
}

// Track creates a "track" node
func Track(children ...Node) Node {
	return &ElementNode{Tag: TagTrack, Children: children, AllowOmission: true}
}

// Video creates a "video" node
func Video(children ...Node) Node {
	return &ElementNode{Tag: TagVideo, Children: children}
}

// Embedded content

// Embed creates an "embed" node
func Embed(children ...Node) Node {
	return &ElementNode{Tag: TagEmbed, Children: children, AllowOmission: true}
}

// Fencedframe creates a "fencedframe" node
func Fencedframe(children ...Node) Node {
	return &ElementNode{Tag: TagFencedFrame, Children: children}
}

// Iframe creates an "iframe" node
func Iframe(children ...Node) Node {
	return &ElementNode{Tag: TagIframe, Children: children}
}

// Object creates an "object" node
func Object(children ...Node) Node {
	return &ElementNode{Tag: TagObject, Children: children}
}

// Picture creates a "picture" node
func Picture(children ...Node) Node {
	return &ElementNode{Tag: TagPicture, Children: children}
}

// Portal creates a "portal" node
func Portal(children ...Node) Node {
	return &ElementNode{Tag: TagPortal, Children: children}
}

// Source creates a "source" node
func Source(children ...Node) Node {
	return &ElementNode{Tag: TagSource, Children: children, AllowOmission: true}
}

// SVG and MathMl

// Svg creates an "svg" node
func Svg(children ...Node) Node {
	return &ElementNode{Tag: TagSvg, Children: children}
}

// Math creates a "math" node
func Math(children ...Node) Node {
	return &ElementNode{Tag: TagMath, Children: children}
}

// Scripting

// Canvas creates a "canvas" node
func Canvas(children ...Node) Node {
	return &ElementNode{Tag: TagCanvas, Children: children}
}

// Noscript creates a "noscript" node
func Noscript(children ...Node) Node {
	return &ElementNode{Tag: TagNoscript, Children: children}
}

// Script creates a "script" node
func Script(children ...Node) Node {
	return &ElementNode{Tag: TagScript, Children: children}
}

// Demarcating edits

// Del creates a "del" node
func Del(children ...Node) Node {
	return &ElementNode{Tag: TagDel, Children: children}
}

// Ins creates a "ins" node
func Ins(children ...Node) Node {
	return &ElementNode{Tag: TagIns, Children: children}
}

// Table content

// Caption creates a "caption" node
func Caption(children ...Node) Node {
	return &ElementNode{Tag: TagCaption, Children: children}
}

// Col creates a "col" node
func Col(children ...Node) Node {
	return &ElementNode{Tag: TagCol, Children: children, AllowOmission: true}
}

// Colgroup creates a "colgroup" node
func Colgroup(children ...Node) Node {
	return &ElementNode{Tag: TagColgroup, Children: children}
}

// Table creates a "table" node
func Table(children ...Node) Node {
	return &ElementNode{Tag: TagTable, Children: children}
}

// Tbody creates a "tbody" node
func Tbody(children ...Node) Node {
	return &ElementNode{Tag: TagTbody, Children: children}
}

// Td creates a "td" node
func Td(children ...Node) Node {
	return &ElementNode{Tag: TagTd, Children: children}
}

// Tfoot creates a "tfoot" node
func Tfoot(children ...Node) Node {
	return &ElementNode{Tag: TagTfoot, Children: children}
}

// Th creates a "th" node
func Th(children ...Node) Node {
	return &ElementNode{Tag: TagTh, Children: children}
}

// Thead creates a "thead" node
func Thead(children ...Node) Node {
	return &ElementNode{Tag: TagThead, Children: children}
}

// Tr creates a "tr" node
func Tr(children ...Node) Node {
	return &ElementNode{Tag: TagTr, Children: children}
}

// Forms

// Button creates a "button" node
func Button(children ...Node) Node {
	return &ElementNode{Tag: TagButton, Children: children}
}

// Datalist creates a "datalist" node
func Datalist(children ...Node) Node {
	return &ElementNode{Tag: TagDatalist, Children: children}
}

// Fieldset creates a "fieldset" node
func Fieldset(children ...Node) Node {
	return &ElementNode{Tag: TagFieldset, Children: children}
}

// Form creates a "form" node
func Form(children ...Node) Node {
	return &ElementNode{Tag: TagForm, Children: children}
}

// Input creates an "input" node
func Input(children ...Node) Node {
	return &ElementNode{Tag: TagInput, Children: children, AllowOmission: true}
}

// Label creates a "label" node
func Label(children ...Node) Node {
	return &ElementNode{Tag: TagLabel, Children: children}
}

// Legend creates a "legend" node
func Legend(children ...Node) Node {
	return &ElementNode{Tag: TagLegend, Children: children}
}

// Meter creates a "meter" node
func Meter(children ...Node) Node {
	return &ElementNode{Tag: TagMeter, Children: children}
}

// Optgroup creates a "optgroup" node
func Optgroup(children ...Node) Node {
	return &ElementNode{Tag: TagOptgroup, Children: children}
}

// Option creates a "option" node
func Option(children ...Node) Node {
	return &ElementNode{Tag: TagOption, Children: children}
}

// Output creates a "output" node
func Output(children ...Node) Node {
	return &ElementNode{Tag: TagOutput, Children: children}
}

// Progress creates a "progress" node
func Progress(children ...Node) Node {
	return &ElementNode{Tag: TagProgress, Children: children}
}

// Select creates a "select" node
func Select(children ...Node) Node {
	return &ElementNode{Tag: TagSelect, Children: children}
}

// Textarea creates a "textarea" node
func Textarea(children ...Node) Node {
	return &ElementNode{Tag: TagTextarea, Children: children}
}

// Interactive elements

// Details creates a "details" node
func Details(children ...Node) Node {
	return &ElementNode{Tag: TagDetails, Children: children}
}

// Dialog creates a "dialog" node
func Dialog(children ...Node) Node {
	return &ElementNode{Tag: TagDialog, Children: children}
}

// Summary creates a "summary" node
func Summary(children ...Node) Node {
	return &ElementNode{Tag: TagSummary, Children: children}
}

// Web Components

// Slot creates a "slot" node
func Slot(children ...Node) Node {
	return &ElementNode{Tag: TagSlot, Children: children}
}

// Template creates a "template" node
func Template(children ...Node) Node {
	return &ElementNode{Tag: TagTemplate, Children: children}
}

//
// Attributes
//

// Accept creates an "accept" attribute
func Accept(accept string) Node {
	return &AttributeNode{Name: AttrAccept, Value: accept}
}

// Action creates an "action" attribute
func Action(action string) Node {
	return &AttributeNode{Name: AttrAction, Value: action}
}

// AutoComplete creates an "autocomplete" attribute
func AutoComplete(autocomplete string) Node {
	return &AttributeNode{Name: AttrAutocomplete, Value: autocomplete}
}

// Charset creates a "charset" attribute
func Charset(charset string) Node {
	return &AttributeNode{Name: AttrCharset, Value: charset}
}

// Checked creates a "checked" attribute
func Checked(checked ...bool) Node {
	if len(checked) > 0 && !checked[0] {
		return nil
	}
	return &AttributeNode{Name: AttrChecked, Value: ""}
}

// Class creates a "class" attribute
func Class(value string) Node {
	return &AttributeNode{Name: AttrClass, Value: value}
}

// Content creates a "content" attribute
func Content(value string) Node {
	return &AttributeNode{Name: AttrContent, Value: value}
}

// Defer creates a "defer" attribute
func Defer(d ...bool) Node {
	if len(d) > 0 && !d[0] {
		return nil
	}
	return &AttributeNode{Name: AttrDefer, Value: ""}
}

// Disabled creates a "disabled" attribute
func Disabled(disabled ...bool) Node {
	if len(disabled) > 0 && !disabled[0] {
		return nil
	}
	return &AttributeNode{Name: AttrDisabled, Value: ""}
}

// Draggable creates a "draggable" attribute
func Draggable(draggable ...bool) Node {
	if len(draggable) > 0 && !draggable[0] {
		return nil
	}
	return &AttributeNode{Name: AttrDraggable, Value: ""}
}

// For creates a "for" attribute
func For(id string) Node {
	return &AttributeNode{Name: AttrFor, Value: id}
}

// Hidden creates a "hidden" attribute
func Hidden(hidden ...bool) Node {
	if len(hidden) > 0 && !hidden[0] {
		return nil
	}
	return &AttributeNode{Name: AttrHidden, Value: ""}
}

// Lang creates a "lang" attribute
func Lang(lang string) Node {
	return &AttributeNode{Name: AttrLang, Value: lang}
}

// Placeholder creates a "placeholder" attribute
func Placeholder(placeholder string) Node {
	return &AttributeNode{Name: AttrPlaceholder, Value: placeholder}
}

// Readonly creates a "readonly" attribute
func Readonly(readonly ...bool) Node {
	if len(readonly) > 0 && !readonly[0] {
		return nil
	}
	return &AttributeNode{Name: AttrReadonly, Value: ""}
}

// Required creates a "required" attribute
func Required(required ...bool) Node {
	if len(required) > 0 && !required[0] {
		return nil
	}
	return &AttributeNode{Name: AttrRequired, Value: ""}
}

// Role creates a "role" attribute
func Role(role string) Node {
	return &AttributeNode{Name: AttrRole, Value: role}
}

// Tabindex creates a "tabindex" attribute
func TabIndex(tabindex int) Node {
	return &AttributeNode{Name: AttrTabIndex, Value: strconv.Itoa(tabindex)}
}

// Target creates a "target" attribute
func Target(target string) Node {
	return &AttributeNode{Name: AttrTarget, Value: target}
}

var targetBlankNode = &AttributeNode{Name: AttrTarget, Value: "_blank"}

// TargetBlank creates a target="_blank" attribute
func TargetBlank() Node {
	return targetBlankNode
}

// TestId creates a "data-testid" attribute
func TestId(value string) Node {
	return &AttributeNode{Name: AttrDatatestid, Value: value}
}

// Id creates an "id" attribute
func Id(value string) Node {
	return &AttributeNode{Name: AttrId, Value: value}
}

// Href creates an "href" attribute
func Href(value string) Node {
	return &AttributeNode{Name: AttrHref, Value: value}
}

// Name creates a "name" attribute
func Name(value string) Node {
	return &AttributeNode{Name: AttrName, Value: value}
}

// Property creates a "property" attribute
func Property(value string) Node {
	return &AttributeNode{Name: AttrProperty, Value: value}
}

// Rel creates a "rel" attribute
func Rel(value string) Node {
	return &AttributeNode{Name: AttrRel, Value: value}
}

var relStylesheetnode = &AttributeNode{Name: AttrRel, Value: "stylesheet"}

// RelStylesheet creates a rel="stylesheet" attribute
func RelStylesheet() Node {
	return relStylesheetnode
}

// Src creates a "src" attribure
func Src(value string) Node {
	return &AttributeNode{Name: AttrSrc, Value: value}
}

// StyleAttr creates a "style" attribute
func StyleAttr(style string) Node {
	return &AttributeNode{Name: AttrStyle, Value: style}
}

// Type creates a "type" attribute
func Type(value string) Node {
	return &AttributeNode{Name: AttrType, Value: value}
}

// Value creates a "value" attribute
func Value(value string) Node {
	return &AttributeNode{Name: AttrValue, Value: value}
}

//
// Utilities
//

// MetaCharset creates a <meta charset="value"> node
func MetaCharset(value string) Node {
	return Meta(Attr(AttrCharset, value))
}

// MetaDescription creates a <meta name="description" content="value"> node
func MetaDescription(value string) Node {
	return Meta(Name("description"), Content(value))
}

// MetaKeywords creates a <meta name="keywords" content="value"> node
func MetaKeywords(value string) Node {
	return Meta(Name("keywords"), Content(value))
}

// MetaViewport creates a <meta name="viewport" content="value" > node
func MetaViewport(value string) Node {
	return Meta(Name("viewport"), Content(value))
}
