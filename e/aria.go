// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

const (
	AttrAriaHidden = "aria-hidden"
)

// AriaHidden returns an aria-hidden attribute
func AriaHidden(hidden ...bool) Node {
	v := ValTrue
	if len(hidden) > 0 && !hidden[0] {
		v = ValFalse
	}
	return &AttributeNode{Name: AttrAriaHidden, Value: v}
}
