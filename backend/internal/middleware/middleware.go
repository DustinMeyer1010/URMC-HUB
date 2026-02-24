package middleware

import "net/http"

type Middleware []func(http.Handler) http.Handler

// Runs the middleware in order
func (m Middleware) MiddlewareChain(h http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}
