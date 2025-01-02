// Copyright 2025 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package e

const (
	PropOGTitle = "og:title"
	PropOGType  = "og:type"
	PropOGUrl   = "og:url"
	PropOGImage = "og:image"

	PropOGAudio           = "og:audio"
	PropOGDescription     = "og:description"
	PropOGDeterminer      = "og:determiner"
	PropOGLocale          = "og:locale"
	PropOGLocaleAlternate = "og:locale:alternate"
	PropOGSiteName        = "og:site_name"
	PropOGVideo           = "og:video"

	PropOGImageUrl       = "og:image:url"
	PropOGImageSecureUrl = "og:image:secure_url"
	PropOGImageType      = "og:image:type"
	PropOGImageWidth     = "og:image:width"
	PropOGImageHeight    = "og:image:height"
	PropOGImageAlt       = "og:image:alt"

	PropOGVideoUrl       = "og:video:url"
	PropOGVideoSecureUrl = "og:video:secure_url"
	PropOGVideoType      = "og:video:type"
	PropOGVideoWidth     = "og:video:width"
	PropOGVideoHeight    = "og:video:height"
	PropOGVideoAlt       = "og:video:alt"

	PropOGAudioSecureUrl = "og:audio:secure_url"
	PropOGAudioType      = "og:audio:type"
)

//
// Basic Metadata
//

// OGTitle creates an og:title meta tag
func OGTitle(title string) Node {
	return Meta(Property(PropOGTitle), Content(title))
}

// OGType creates an og:type meta tag
func OGType(value string) Node {
	return Meta(Property(PropOGType), Content(value))
}

// OGUrl creates an og:url meta tag
func OGUrl(url string) Node {
	return Meta(Property(PropOGUrl), Content(url))
}

//
// Optional Metadata
//

// OGImage creates an og:image meta tag
func OGImage(value string) Node {
	return Meta(Property(PropOGImage), Content(value))
}

// OGAudio creates an og:audio meta tag
func OGAudio(value string) Node {
	return Meta(Property(PropOGAudio), Content(value))
}

// OGDescription creates an og:description meta tag
func OGDescription(description string) Node {
	return Meta(Property(PropOGDescription), Content(description))
}

// OGDeterminer creates an og:determiner meta tag
func OGDeterminer(value string) Node {
	return Meta(Property(PropOGDeterminer), Content(value))
}

// OGLocale creates an og:locale meta tag
func OGLocale(locale string) Node {
	return Meta(Property(PropOGLocale), Content(locale))
}

// OGLocaleAlternate creates an og:locale:alternate meta tag
func OGLocaleAlternate(value string) Node {
	return Meta(Property(PropOGLocaleAlternate), Content(value))
}

// OGSiteName creates an og:site_name meta tag
func OGSiteName(name string) Node {
	return Meta(Property(PropOGSiteName), Content(name))
}

// OGVideo creates an og:video meta tag
func OGVideo(value string) Node {
	return Meta(Property(PropOGVideo), Content(value))
}

//
// Structured Properties
//

// OGImageUrl creates an og:image:url meta tag
func OGImageUrl(url string) Node {
	return Meta(Property(PropOGImageUrl), Content(url))
}

// OGImageSecureUrl creates an og:image:secure_url meta tag
func OGImageSecureUrl(url string) Node {
	return Meta(Property(PropOGImageSecureUrl), Content(url))
}

// OGImageType creates an og:image:type meta tag
func OGImageType(value string) Node {
	return Meta(Property(PropOGImageType), Content(value))
}

// OGImageWidth creates an og:image:width meta tag
func OGImageWidth(width string) Node {
	return Meta(Property(PropOGImageWidth), Content(width))
}

// OGImageHeight creates an og:image:height meta tag
func OGImageHeight(height string) Node {
	return Meta(Property(PropOGImageHeight), Content(height))
}

// OGImageAlt creates an og:image:alt meta tag
func OGImageAlt(alt string) Node {
	return Meta(Property(PropOGImageAlt), Content(alt))
}

// OGVideoUrl creates an og:video:url meta tag
func OGVideoUrl(vurllue string) Node {
	return Meta(Property(PropOGVideoUrl), Content(vurllue))
}

// OGVideoSecureUrl creates an og:video:secure_url meta tag
func OGVideoSecureUrl(url string) Node {
	return Meta(Property(PropOGVideoSecureUrl), Content(url))
}

// OGVideoType creates an og:video:type meta tag
func OGVideoType(value string) Node {
	return Meta(Property(PropOGVideoType), Content(value))
}

// OGVideoWidth creates an og:video:width meta tag
func OGVideoWidth(width string) Node {
	return Meta(Property(PropOGVideoWidth), Content(width))
}

// OGVideoHeight creates an og:video:height meta tag
func OGVideoHeight(height string) Node {
	return Meta(Property(PropOGVideoHeight), Content(height))
}

// OGVideoAlt creates an og:video:alt meta tag
func OGVideoAlt(alt string) Node {
	return Meta(Property(PropOGVideoAlt), Content(alt))
}

// OGAudioSecureUrl creates an og:audio:secure_url meta tag
func OGAudioSecureUrl(url string) Node {
	return Meta(Property(PropOGAudioSecureUrl), Content(url))
}

// OGAudioType creates an og:audio:type meta tag
func OGAudioType(value string) Node {
	return Meta(Property(PropOGAudioType), Content(value))
}
