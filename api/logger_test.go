package api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Handler Mock

type MockedHandler struct {
	http.Handler
}

func (h *MockedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Method = "GET"
	r.RequestURI = "/test"
}

func TestLogger(t *testing.T) {
	// change buffer so that we can check
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	r, _ := http.NewRequest("GET", "http://example.com/test", nil)
	w := httptest.NewRecorder()

	mockedHandler := new(MockedHandler)
	handler := Logger(mockedHandler, "testroute")
	handler.ServeHTTP(w, r)

	logged := buf.String()
	assert.NotEqual(t, "", logged, "Unexpected logged line"+logged)
	assert.Contains(t, logged, "GET\t/test\ttestroute", "Unexpected logged line")
}
