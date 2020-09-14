package handler

import (
	"log"
	"net/http"
)

var (
	_ http.Handler = (*LoggingHandler)(nil)
)

// LoggingHandler is a logging handler
type LoggingHandler struct {
	logger *log.Logger
}

// NewHandler creates and returns a logging handler
func NewLoggingHandler(enable bool, logger *log.Logger) http.Handler {
	if !enable {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {})
	}
	return &LoggingHandler{logger}
}

// ServeHTTP serves while writing access log
func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch w.(type) {
	case *Response:
		resp := w.(*Response)
		h.logger.Printf(
			`%s %s %s %s %d "%s" "%s"`,
			req.RemoteAddr,
			req.Method,
			req.RequestURI,
			req.Proto,
			resp.statusCode,
			req.Referer(),
			req.UserAgent(),
		)
	default:
	}
}
