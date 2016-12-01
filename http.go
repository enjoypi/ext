package ext

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

func PerformHttpRequest(r http.Handler, method, path string, header http.Header, reqBody io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, reqBody)
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
	if err == nil && len(m) > 0 {
		return m
	}
	return nil
}

func ReadMockFile(filepath string) string {
	file, err := os.Open(filepath) // For read access.
	if err != nil {
		return ""
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	sdata := string(data)
	return sdata
}
