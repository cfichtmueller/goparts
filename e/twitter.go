// Copyright 2025 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

const (
	PropTwitterCard        = "twitter:card"
	PropTwitterSite        = "twitter:site"
	PropTwitterSiteId      = "twitter:site:id"
	PropTwitterCreator     = "twitter:creator"
	PropTwitterCreatorId   = "twitter:creator:id"
	PropTwitterDescription = "twitter:description"
	PropTwitterTitle       = "twitter:title"
	PropTwitterImage       = "twitter:image"
	PropTwitterImageAlt    = "twitter:image:alt"
)

// TwitterCard creates a twitter:card meta tag
func TwitterCard(value string) Node {
	return Meta(Property(PropTwitterCard), Content(value))
}

// TwitterSite creates a twitter:site meta tag
func TwitterSite(value string) Node {
	return Meta(Property(PropTwitterSite), Content(value))
}

// TwitterSiteId creates a twitter:site:id meta tag
func TwitterSiteId(value string) Node {
	return Meta(Property(PropTwitterSiteId), Content(value))
}

// TwitterCreator creates a twitter:creator meta tag
func TwitterCreator(value string) Node {
	return Meta(Property(PropTwitterCreator), Content(value))
}

// TwitterCreatorId creates a twitter:creator:id meta tag
func TwitterCreatorId(value string) Node {
	return Meta(Property(PropTwitterCreatorId), Content(value))
}

// TwitterDescription creates a twitter:description meta tag
func TwitterDescription(value string) Node {
	return Meta(Property(PropTwitterDescription), Content(value))
}

// TwitterTitle creates a twitter:title meta tag
func TwitterTitle(value string) Node {
	return Meta(Property(PropTwitterTitle), Content(value))
}

// TwitterImage creates a twitter:image meta tag
func TwitterImage(value string) Node {
	return Meta(Property(PropTwitterImage), Content(value))
}

// TwitterImageAlt creates a twitter:image:alt meta tag
func TwitterImageAlt(value string) Node {
	return Meta(Property(PropTwitterImageAlt), Content(value))
}
