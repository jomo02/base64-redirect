package base64redirect

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	fmt.Println("Registering Base64Redirect module")
	caddy.RegisterModule(Base64Redirect{})
}

// Base64Redirect is a Caddy HTTP handler module that encodes the original URL and redirects.
type Base64Redirect struct {
	Target string `json:"target,omitempty"` // Target base URL to redirect to
}

// CaddyModule returns the Caddy module information.
func (Base64Redirect) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.base64_redirect",
		New: func() caddy.Module { return new(Base64Redirect) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (br Base64Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Printf("Handling request for URL: %s\n", r.URL.String())

	// Check if target was set
	if br.Target == "" {
		http.Error(w, "Target URL is not set", http.StatusInternalServerError)
		return nil
	}

	// Check that target is valid (at least correct protocol)
	if !strings.HasPrefix(br.Target, "http://") && !strings.HasPrefix(br.Target, "https://") {
		http.Error(w, "Invalid Target URL", http.StatusInternalServerError)
		return nil
	}

	// Encode the original URL path and query to Base64
	originalURL := r.URL.String()
	encodedURL := base64.URLEncoding.EncodeToString([]byte(originalURL))

	// Construct the redirect URL
	redirectURL := fmt.Sprintf("%s%s", br.Target, encodedURL)

	// Send the redirect response
	http.Redirect(w, r, redirectURL, http.StatusFound)
	return nil
}

// Interface guards
var (
	_ caddy.Module                = (*Base64Redirect)(nil)
	_ caddyhttp.MiddlewareHandler = (*Base64Redirect)(nil)
)