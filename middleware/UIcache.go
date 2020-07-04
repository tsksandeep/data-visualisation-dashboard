package middleware

import (
	"net/http"
	"regexp"
)

const (
	cacheControl        = "Cache-Control"
	cacheControlNoStore = "no-store"
	cacheControlPublic  = "public, max-age=604800, immutable"
)

var (
	noCachePaths = []string{"/index.html"}
	cachePaths   = regexp.MustCompile(`^/.*\.(css|js|json|png|jpg|jpeg|ico|svg|ttf|woff|woff2)$`)
)

func UICacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheControlValue := cacheControlNoStore
		if shouldCache(r.URL.Path) {
			cacheControlValue = cacheControlPublic
		}
		w.Header().Set(cacheControl, cacheControlValue)
		next.ServeHTTP(w, r)
	})
}

func shouldCache(path string) bool {
	for _, f := range noCachePaths {
		if path == f {
			return false
		}
	}
	return cachePaths.MatchString(path)
}
