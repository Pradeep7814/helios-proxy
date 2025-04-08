package unit

import (
	"helios-proxy/pkg/proxy"
	"net/http/httptest"
	"testing"
)

func TestProxyHandler(t *testing.T) {
	handler := proxy.NewHandler()

	req := httptest.NewRequest("GET", "/some-endpoint", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	expected := "Proxy handler response"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
