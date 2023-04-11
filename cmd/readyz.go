package main

import (
	"net/http"
	"sync/atomic"
)

// readyz is a readiness probe.
func readyz(isReady *atomic.Value) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isReady == nil || !isReady.Load().(bool) {
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}