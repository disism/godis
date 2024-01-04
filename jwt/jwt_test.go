package jwt

import (
	"testing"
	"time"
)

func TestJWT_GenerateAndValidate(t *testing.T) {
	secret := "test_secret"

	jwtToken := New(SigningMethodHS256, WithExpiresAt(time.Now().Add(time.Hour)))

	tokenString, err := jwtToken.Generate(secret)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	t.Logf("Generated token: %v", tokenString)

	// Validate the token
	claims, err := jwtToken.Validate(tokenString, secret)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	t.Logf("Validated token: %v", claims)
}

func TestJWT_ValidateExpiredToken(t *testing.T) {
	secret := "test_secret"
	expiredTime := time.Now().Add(-time.Hour)

	jwtToken := New(SigningMethodHS256, WithExpiresAt(expiredTime))

	tokenString, err := jwtToken.Generate(secret)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	_, err = jwtToken.Validate(tokenString, secret)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}
