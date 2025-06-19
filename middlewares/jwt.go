package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/internals/common/response"
	"github.com/shehbaazsk/go-commerce/internals/constants"
	"github.com/shehbaazsk/go-commerce/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Missing Authorization header", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(c, http.StatusUnauthorized, "Invalid token format", nil)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token", err)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*utils.CustomClaims)
		if !ok || claims.TokenType != "access" {
			response.Error(c, http.StatusUnauthorized, "Invalid token claims", nil)
			c.Abort()
			return
		}

		// // Set claims into context
		c.Set(string(constants.UserIDKey), claims.UserID)
		ctx := context.WithValue(c.Request.Context(), constants.UserIDKey, claims.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
