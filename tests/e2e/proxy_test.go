package e2e

import (
	"helios-proxy/pkg/auth"
	"helios-proxy/pkg/proxy"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyAuthorization(t *testing.T) {
	// Create a new request to test the proxy
	req, err := http.NewRequest("GET", "/some-endpoint", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Add authorization header (JWT token for example)
	token, _ := auth.GenerateToken("testuser") // Assuming GenerateToken is a function in auth package
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a proxy handler
	handler := http.HandlerFunc(proxy.HandleRequest)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Further assertions can be added here to check response body, headers, etc.
}

func TestProxyRouting(t *testing.T) {
	router := proxy.NewRouter()

	// Test undefined route
	req := httptest.NewRequest("GET", "/undefined-route", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusNotFound)
	}
}
