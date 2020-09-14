package handler

import (
	"net/http"
)

var (
	_ http.ResponseWriter = (*Response)(nil)
)

// ResponseWriter is a response writer that captures response status code
type Response struct {
	http.ResponseWriter
	statusCode int
}

// NewResponseWriter creates a responseWriter with default status code 200
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w, http.StatusOK}
}

// WriteHeader captures the status code to return
func (w *Response) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// Middleware is a chain of response handlers
type Middleware struct {
	handlers []http.Handler
}

// NewMiddleware creates a Middleware instance
func NewMiddleware(handlers []http.Handler) http.Handler {
	return &Middleware{handlers}
}

// ServeHTTP executes a chain of response handlers
func (h *Middleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp := NewResponse(w)
	for _, hdr := range h.handlers {
		hdr.ServeHTTP(resp, req)
	}
}
