package security

import (
	"net/http"
	"strings"
)

// LocalOnlyMiddleware ensures only local requests can access sensitive AI routes
func LocalOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		
		// Only allow localhost or the internal network bridge
		if !strings.HasPrefix(remoteAddr, "127.0.0.1") && !strings.HasPrefix(remoteAddr, "[::1]") {
			http.Error(w, "Access Denied: External AI execution is forbidden.", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
