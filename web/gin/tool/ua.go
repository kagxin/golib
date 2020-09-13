package tool

import (
	"io"
	"net/http"
	"net/http/httptest"
)

type header struct {
	Key   string
	Value string
}

// PerformRequest for ua test
func PerformRequest(r http.Handler, method, path string, body io.Reader, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
