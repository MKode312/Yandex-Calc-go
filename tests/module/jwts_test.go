package module_test

import (
	"calculator_go/internal/utils/orchestrator/jwts"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyJWTToken(t *testing.T) {
	userID := int64(12345)
	token, err := jwts.GenerateJWTToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	tests := []struct {
		name       string
		token      string
		expectUser string
		expectErr  bool
	}{
		{
			name:       "valid token",
			token:      token,
			expectUser: "12345",
			expectErr:  false,
		},
        {
            name:       "invalid token",
            token:      "invalid.token.string",
            expectUser: "",
            expectErr:  true,
        },
        {
            name:       "expired token",
            token:      createExpiredToken(userID),
            expectUser: "",
            expectErr:  true,
        },
    }

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            userID, err := jwts.VerifyJWTToken(tt.token)

            if (err != nil) != tt.expectErr {
                t.Errorf("expected error status %v, got %v", tt.expectErr, err != nil)
            }

            if userID != tt.expectUser {
                t.Errorf("expected userID %s, got %s", tt.expectUser, userID)
            }
        })
    }
}

// createExpiredToken creates a JWT token that is already expired for testing purposes.
func createExpiredToken(userID int64) string {
	os.Setenv("JWT_SECRET_KEY", "test_secret")
	secretKey := os.Getenv("JWT_SECRET_KEY")
	now := time.Now()
	userIDStr := fmt.Sprintf("%d", userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userid": userIDStr,
        "iat": now.Add(-time.Hour).Unix(), // issued at one hour ago
        "nbf": now.Add(-time.Hour).Unix(), // not before one hour ago
        "exp": now.Add(-time.Minute).Unix(), // expired one minute ago
    })

	tokenString, _ := token.SignedString([]byte(secretKey))
	return tokenString
}