package unit

import (
	"helios-proxy/pkg/auth"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set the JWT_SECRET environment variable for tests
	os.Setenv("JWT_SECRET", "test_secret_key")
	code := m.Run()
	os.Exit(code)
}

func TestGenerateToken(t *testing.T) {
	token, err := auth.GenerateToken("testuser")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("Generated token is empty")
	}
}
