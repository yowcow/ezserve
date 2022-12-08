package cors

import "net/http"

func NewHandler(handler http.Handler, allow bool) http.Handler {
	return &Handler{
		next:  handler,
		allow: allow,
	}
}

type Handler struct {
	next  http.Handler
	allow bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h.allow {
		if origin := req.Header.Get("origin"); origin != "" {
			w.Header().Add("access-control-allow-origin", origin)
			w.Header().Add("access-control-allow-credentials", "true")
		}
	}
	h.next.ServeHTTP(w, req)
}
