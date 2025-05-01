// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

import (
	"io"
	"text/template"
)

const (
	ValFalse = "false"
	ValTrue  = "true"
)

var (
	space   = []byte(" ")
	newline = []byte("\n")
)

type Node interface {
	Render(w io.Writer) error
	Write(w io.Writer) (int, error)
}

type ElementNode struct {
	Tag           string
	Children      []Node
	AllowOmission bool
}

// E creates a new element
func E(tag string, children ...Node) Node {
	return &ElementNode{
		Tag:      tag,
		Children: children,
	}
}

func (e *ElementNode) Render(w io.Writer) error {
	_, err := e.Write(w)
	return err
}

func (e *ElementNode) Write(w io.Writer) (int, error) {
	n := 0
	cn, err := w.Write([]byte("<" + e.Tag))
	if err != nil {
		return n + cn, err
	}
	n += cn
	childCount := 0
	for _, c := range e.Children {
		if a, ok := c.(*AttributeNode); ok {
			childCount++
			cn, err := w.Write(space)
			if err != nil {
				return n + cn, err
			}
			n += cn
			cn, err = a.Write(w)
			if err != nil {
				return n + cn, err
			}
			n += cn
		}
	}

	if e.AllowOmission && childCount == len(e.Children) {
		cn, err := w.Write([]byte(" />"))
		return n + cn, err
	}

	cn, err = w.Write([]byte(">"))
	if err != nil {
		return n + cn, err
	}
	n += cn
	hasSub := false
	for _, c := range e.Children {
		if c == nil {
			continue
		}
		if _, ok := c.(*AttributeNode); ok {
			continue
		}
		if _, ok := c.(*ElementNode); ok {
			hasSub = true
			cn, err = w.Write(newline)
			if err != nil {
				return n + cn, err
			}
			n += cn
		}
		cn, err = c.Write(w)
		if err != nil {
			return n + cn, err
		}
		n += cn
	}
	if hasSub {
		cn, err = w.Write(newline)
		if err != nil {
			return n + cn, err
		}
		n += cn
	}
	cn, err = w.Write([]byte("</" + e.Tag + ">"))
	return n + cn, err
}

type AttributeNode struct {
	Name  string
	Value string
}

// Attr creates a new attribute
func Attr(name, value string) Node {
	return &AttributeNode{
		Name:  name,
		Value: value,
	}
}

func (a *AttributeNode) Render(w io.Writer) error {
	_, err := w.Write([]byte(a.Name))
	return err
}

func (a *AttributeNode) Write(w io.Writer) (int, error) {
	if a.Value == "" {
		return w.Write([]byte(a.Name))
	}
	return w.Write([]byte(a.Name + "=\"" + a.Value + "\""))
}

type RawNode struct {
	Bytes []byte
}

func Raw(text string) Node {
	return &RawNode{Bytes: []byte(text)}
}

func (n *RawNode) Render(w io.Writer) error {
	_, err := w.Write(n.Bytes)
	return err
}

func (n *RawNode) Write(w io.Writer) (int, error) {
	return w.Write(n.Bytes)
}

type TextNode struct {
	Value string
}

// Text creates a new text node
func Text(value string) Node {
	return &TextNode{Value: value}
}

func (t *TextNode) Render(w io.Writer) error {
	_, err := t.Write(w)
	return err
}

func (t *TextNode) Write(w io.Writer) (int, error) {
	return w.Write([]byte(template.HTMLEscapeString(t.Value)))
}

type GroupNode struct {
	children []Node
}

// Group creates a group of nodes
func Group(children ...Node) Node {
	return &GroupNode{children: children}
}

func (g *GroupNode) Render(w io.Writer) error {
	_, err := g.Write(w)
	return err
}

func (g *GroupNode) Write(w io.Writer) (int, error) {
	n := 0
	first := true
	for _, c := range g.children {
		if c == nil {
			continue
		}
		if first {
			first = false
		} else {
			cn, err := w.Write(newline)
			if err != nil {
				return n + cn, err
			}
			n += cn
		}
		cn, err := c.Write(w)
		if err != nil {
			return n + cn, err
		}
		n += cn
	}
	return n, nil
}

// Functions

// If conditionally renders a node
func If(condition bool, value Node) Node {
	if condition {
		return value
	}
	return nil
}

// Iff conditionally creates a node
func Iff(condition bool, f func() Node) Node {
	if condition {
		return f()
	}
	return nil
}

// Mapf maps a set of values to nodes
func Mapf[T any](values []T, f func(t T) Node) Node {
	children := make([]Node, len(values))
	for i, v := range values {
		children[i] = f(v)
	}
	return &GroupNode{children: children}
}

func F[T any](fg func(T) Node, t T) func() Node {
	return func() Node {
		return fg(t)
	}
}
