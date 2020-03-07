package logging

import (
	"log"
	"net/http"
)

var (
	_ http.ResponseWriter = (*responseWriter)(nil)
	_ http.Handler        = (*Handler)(nil)
)

// responseWriter is a response writer that captures response status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newResponseWriter creates a responseWriter with default status code 200
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

// WriteHeader captures the status code to return
func (w *responseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// Handler is a logging handler
type Handler struct {
	handler http.Handler
	logger  *log.Logger
}

// NewHandler creates and returns a logging handler
func NewHandler(h http.Handler, l *log.Logger) *Handler {
	return &Handler{h, l}
}

// ServeHTTP serves while writing access log
func (l *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp := newResponseWriter(w)
	l.handler.ServeHTTP(resp, req)
	l.logger.Printf(
		`%s %s %s %s %d "%s" "%s"`,
		req.RemoteAddr,
		req.Method,
		req.RequestURI,
		req.Proto,
		resp.statusCode,
		req.Referer(),
		req.UserAgent(),
	)
	return
}
