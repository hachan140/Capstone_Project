package middleware

import (
	"be-capstone-project/src/internal/core/web/context"
	"net/http"
)

// AdvancedResponseWriter middleware responsible to replace
// default response writer with a AdvancedResponseWriter
func AdvancedResponseWriter() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(context.NewAdvancedResponseWriter(w), r)
		})
	}
}
