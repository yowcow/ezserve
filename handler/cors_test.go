package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSHandler(t *testing.T) {
	cases := []struct {
		title       string
		allowOrigin string
	}{
		{
			"w/o allowOrigin",
			"",
		},
		{
			"w/ allowOrigin",
			"https://example.com",
		},
	}

	var client http.Client

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			h := NewMiddleware([]http.Handler{
				NewCORSHandler(c.allowOrigin),
			})
			ts := httptest.NewServer(h)
			defer ts.Close()

			req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
			if err != nil {
				t.Fatal("failed creating a request:", err)
			}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal("failed retrieving a response:", err)
			}
			defer resp.Body.Close()

			actual := resp.Header.Get("access-control-allow-origin")
			if c.allowOrigin != actual {
				t.Error("expected", c.allowOrigin, "but got", actual)
			}
		})
	}
}
