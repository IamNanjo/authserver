package middleware

import "net/http"

// MiddlewareHandler should return true to indicate that we should continue
type middlewareHandler func(http.ResponseWriter, *http.Request) bool

func WithMiddleware(handler http.HandlerFunc, middleware ...middlewareHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range middleware {
			shouldContinue := handler(w, r)
			if !shouldContinue {
				return
			}
		}

		handler(w, r)
	}
}
