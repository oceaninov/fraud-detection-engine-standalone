package hashing

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateJWT generates a new JWT token with the given role, email, userId, secretKey, and expiration time.
func GenerateJWT(role, email, userId, secretKey string, expirationTime time.Duration) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"role":   role,
		"exp":    time.Now().Add(expirationTime).Unix(), // Expiration time
		"iat":    time.Now().Unix(),                     // Issued at time
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// VerifyJWT verifies a JWT token and returns the claims if valid.
func VerifyJWT(tokenString, secretKey string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// Handle parsing errors
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Validate the token and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")
}
