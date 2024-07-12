package handlers

import (
	"net/http"
)

// cacheControlHandler adds Cache-Control headers to static files
func CacheControlHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set cache control headers for 7 days
		w.Header().Set("Cache-Control", "public, max-age=604800")
		h.ServeHTTP(w, r)
	})
}
