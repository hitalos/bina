//go:build dev

package public

import "net/http"

var (
	FS = http.Dir("cmd/public")
)
