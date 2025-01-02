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
	if _, err := w.Write([]byte("<" + e.Tag)); err != nil {
		return err
	}
	childCount := 0
	for _, c := range e.Children {
		if a, ok := c.(*AttributeNode); ok {
			childCount++
			if _, err := w.Write(space); err != nil {
				return err
			}
			if err := a.Render(w); err != nil {
				return err
			}
		}
	}

	if e.AllowOmission && childCount == len(e.Children) {
		_, err := w.Write([]byte(" />"))
		return err
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}
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
			if _, err := w.Write(newline); err != nil {
				return err
			}
		}
		if err := c.Render(w); err != nil {
			return err
		}
	}
	if hasSub {
		if _, err := w.Write(newline); err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("</" + e.Tag + ">"))
	return err
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
	if a.Value == "" {
		_, err := w.Write([]byte(a.Name))
		return err
	}
	_, err := w.Write([]byte(a.Name + "=\"" + a.Value + "\""))
	return err
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

type TextNode struct {
	Value string
}

// Text creates a new text node
func Text(value string) Node {
	return &TextNode{Value: value}
}

func (t *TextNode) Render(w io.Writer) error {
	_, err := w.Write([]byte(template.HTMLEscapeString(t.Value)))
	return err
}

type GroupNode struct {
	children []Node
}

// Group creates a group of nodes
func Group(children ...Node) Node {
	return &GroupNode{children: children}
}

func (g *GroupNode) Render(w io.Writer) error {
	first := true
	for _, c := range g.children {
		if c == nil {
			continue
		}
		if first {
			first = false
		} else {
			w.Write(newline)
		}
		if err := c.Render(w); err != nil {
			return err
		}
	}
	return nil
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
