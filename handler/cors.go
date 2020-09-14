package handler

import (
	"net/http"
)

var (
	_ http.Handler = (*CORSHandler)(nil)
)

type CORSHandler struct {
	allowOrigin string
}

func NewCORSHandler(allowOrigin string) *CORSHandler {
	return &CORSHandler{allowOrigin}
}

func (h *CORSHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h.allowOrigin != "" {
		w.Header().Add("Access-Control-Allow-Origin", h.allowOrigin)
	}
}
