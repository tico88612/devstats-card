package web

import _ "embed"

//go:embed index.html
var indexHTML []byte

// IndexPage returns the landing/preview page served at the site root.
func IndexPage() []byte {
	return indexHTML
}