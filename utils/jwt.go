package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shehbaazsk/go-commerce/config"
)

type CustomClaims struct {
	UserID    uint64 `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint64) (string, error) {
	JWT_EXPIRATION := config.App.JWT_EXPIRATION
	expirationTime := time.Now().Add(time.Duration(JWT_EXPIRATION) * time.Minute)

	claims := &CustomClaims{
		UserID:    userID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign it
	tokenString, err := token.SignedString([]byte(config.App.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID uint64) (string, error) {
	JWT_REFRESH_EXPIRATION := config.App.JWT_REFRESH_EXPIRATION
	expirationTime := time.Now().Add(time.Duration(JWT_REFRESH_EXPIRATION) * time.Minute)

	claims := &CustomClaims{
		UserID:    userID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.App.JWT_SECRET))
}
