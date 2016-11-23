package ext

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

func PerformHttpRequest(r http.Handler, method, path string, header http.Header) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	if header != nil {
		req.Header = header
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func DecodeJson(r io.Reader) map[string]interface{} {
	var m map[string]interface{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&m)
	if err == nil {
		return m
	}
	return nil
}
