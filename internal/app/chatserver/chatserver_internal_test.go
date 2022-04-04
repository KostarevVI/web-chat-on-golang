// Package chatserver as main server app
package chatserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestChatServer_HandlerHello ...
func TestChatServer_HandlerHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handlerHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
