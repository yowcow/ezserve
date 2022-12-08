package cors

import "net/http"

func NewHandler(handler http.Handler, allowOrigin bool) http.Handler {
	return &Handler{
		next:        handler,
		allowOrigin: allowOrigin,
	}
}

type Handler struct {
	next        http.Handler
	allowOrigin bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("origin"); origin != "" {
		w.Header().Add("access-control-allow-origin", origin)
	}
	h.next.ServeHTTP(w, req)
}
