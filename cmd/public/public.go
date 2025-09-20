//go:build !dev

package public

import (
	"embed"
	"net/http"
)

var (
	//go:embed assets favicon.ico images index.html manifest.json pwabuilder-sw.js
	fs embed.FS

	FS = http.FS(fs)
)
