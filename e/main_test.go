// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE fil

package e

import (
	"strings"
	"testing"
)

type nodeTest struct {
	n Node
	e string
}

func TestElementNode(t *testing.T) {
	tests := []nodeTest{
		{n: Br(), e: "<br />"},
		{n: Div(), e: "<div></div>"},
		{n: Div(Class("font-bold"), Role("dialog")), e: "<div class=\"font-bold\" role=\"dialog\"></div>"},
		{n: A(Href("/api/me")), e: "<a href=\"/api/me\"></a>"},
		{n: Div(If(false, Div())), e: "<div></div>"},
		{n: Div(If(true, Div())), e: "<div>\n<div></div>\n</div>"},
		{n: Div(Text("Hello World")), e: "<div>Hello World</div>"},
		{n: Div(Raw("Hello World")), e: "<div>Hello World</div>"},
		{n: Meta(Charset("UTF-8")), e: "<meta charset=\"UTF-8\" />"},
	}

	for i, tst := range tests {
		w := &strings.Builder{}
		if err := tst.n.Render(w); err != nil {
			t.Errorf("unable to render node for test %d: %v", i, err)
		}
		actual := w.String()
		if actual != tst.e {
			t.Errorf("Test %d: got '%s' but expected '%s'", i, actual, tst.e)
		}
	}
}

func TestTextNode(t *testing.T) {
	tests := []nodeTest{
		{n: Text("Hello"), e: "Hello"},
		{n: Text("'ol chap"), e: "&#39;ol chap"},
		{n: Text("this & that"), e: "this &amp; that"},
		{n: Text("<div>foo</div>"), e: "&lt;div&gt;foo&lt;/div&gt;"},
	}

	for _, tst := range tests {
		w := &strings.Builder{}
		if err := tst.n.Render(w); err != nil {
			t.Errorf("unable to render node: %v", err)
		}
		actual := w.String()
		if actual != tst.e {
			t.Errorf("Got '%s' but expected '%s'", actual, tst.e)
		}
	}
}

func TestRawNode(t *testing.T) {
	tests := []nodeTest{
		{n: Raw("Hello"), e: "Hello"},
		{n: Raw("'ol chap"), e: "'ol chap"},
		{n: Raw("this & that"), e: "this & that"},
		{n: Raw("<div>foo</div>"), e: "<div>foo</div>"},
	}

	for _, tst := range tests {
		w := &strings.Builder{}
		if err := tst.n.Render(w); err != nil {
			t.Errorf("unable to render node: %v", err)
		}
		actual := w.String()
		if actual != tst.e {
			t.Errorf("Got '%s' but expected '%s'", actual, tst.e)
		}
	}
}
