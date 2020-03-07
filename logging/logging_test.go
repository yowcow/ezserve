package logging

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestLogging(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	h := NewHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.RequestURI {
		case "/200":
			fmt.Fprint(w, "200 OK")
		case "/500":
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not Found")
		}
	}), logger)

	ts := httptest.NewServer(h)
	defer ts.Close()

	cases := []struct {
		title          string
		uri            string
		referer        string
		expectedStatus int
		expectedBody   string
		expectedLog    string
	}{
		{
			"200 OK",
			"/200",
			"",
			http.StatusOK,
			"200 OK",
			`GET /200 HTTP/1.1 200 "" "Go-http-client/1.1"
`,
		},
		{
			"500 Internal Server Error",
			"/500",
			"http://foobar/fizbuz",
			http.StatusInternalServerError,
			"500 Internal Server Error",
			`GET /500 HTTP/1.1 500 "http://foobar/fizbuz" "Go-http-client/1.1"
`,
		},
		{
			"404 Not Found",
			"/404",
			"http://foobar/fizbuz",
			http.StatusNotFound,
			"404 Not Found",
			`GET /404 HTTP/1.1 404 "http://foobar/fizbuz" "Go-http-client/1.1"
`,
		},
	}

	var client http.Client
	re := regexp.MustCompile(`^127\.0\.0\.1\:\d+\s+`)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, ts.URL+c.uri, nil)
			if err != nil {
				t.Fatal("failed creating a request:", err)
			}
			if c.referer != "" {
				req.Header.Add("Referer", c.referer)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal("failed retrieving a response:", err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("failed reading body:", err)
			}
			defer resp.Body.Close()

			if c.expectedStatus != resp.StatusCode {
				t.Error("expected status", c.expectedStatus, "but got", resp.StatusCode)
			}
			if c.expectedBody != string(body) {
				t.Error("expected body", c.expectedBody, "but got", string(body))
			}

			logOutput := buf.String()
			logOutput = re.ReplaceAllString(logOutput, "")
			buf.Reset()

			if c.expectedLog != logOutput {
				t.Errorf("expected log '%s' but got '%s'", c.expectedLog, logOutput)
			}
		})
	}
}
