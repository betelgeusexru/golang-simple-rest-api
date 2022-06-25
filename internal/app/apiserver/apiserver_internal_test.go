package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	server := New(NewConfig())
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	
	server.handleHello().ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Body.String(), "hello")
}