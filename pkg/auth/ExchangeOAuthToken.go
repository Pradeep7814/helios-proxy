package auth

import (
	"errors"
)

// ExchangeOAuthToken simulates an OAuth token exchange for testing purposes.
func ExchangeOAuthToken(username string) (string, error) {
	if username == "" {
		return "", errors.New("username cannot be empty")
	}
	return GenerateToken(username)
}
