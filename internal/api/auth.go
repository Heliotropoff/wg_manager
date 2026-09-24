package api

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

const apiKeyHeader = "X-API-Key"

// RequireAPIKey protects an HTTP handler with an API key supplied in X-API-Key.
// An empty server-side key keeps the protected handler unavailable.
func RequireAPIKey(expectedKey string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		providedKey := req.Header.Get(apiKeyHeader)

		// The client sends the original API key, not its hash. Hashing both keys
		// produces fixed-length values so they can be compared in constant time,
		// reducing the risk of timing attacks without changing the request format.
		expectedHash := sha256.Sum256([]byte(expectedKey))
		providedHash := sha256.Sum256([]byte(providedKey))
		if expectedKey == "" || providedKey == "" || subtle.ConstantTimeCompare(expectedHash[:], providedHash[:]) != 1 {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, req)
	})
}
