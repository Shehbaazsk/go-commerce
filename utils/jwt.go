package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shehbaazsk/go-commerce/config"
)

// Claims defines the JWT claims structure
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.App.JWT_EXPIRATION_MINUTES) * time.Minute)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.App.JWTSecret))
}

// ValidateToken parses and validates JWT token
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
