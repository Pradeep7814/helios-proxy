package e2e

import (
	"helios-proxy/pkg/auth"
	"helios-proxy/pkg/proxy"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestJWTAuthorization tests that a valid JWT token allows access to the protected endpoint.
func TestJWTAuthorization(t *testing.T) {
	// Create a test server with the proxy handler
	ts := httptest.NewServer(proxy.NewHandler())
	defer ts.Close()

	// Generate a valid JWT token for testing
	token, err := auth.GenerateToken("testuser")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Create a request with the JWT token
	req, err := http.NewRequest("GET", ts.URL+"/protected", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
}

func TestOAuthAuthorization(t *testing.T) {
	// Create a test server with the proxy handler
	ts := httptest.NewServer(proxy.NewHandler())
	defer ts.Close()

	// Simulate an OAuth token exchange and get a valid token
	token, err := auth.ExchangeOAuthToken("testuser")
	if err != nil {
		t.Fatalf("Failed to exchange OAuth token: %v", err)
	}

	// Create a request with the OAuth token
	req, err := http.NewRequest("GET", ts.URL+"/protected", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}
}
