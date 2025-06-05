package middlewares

import (
	"context"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/common/response"
)

// RBAC middleware (role-based access control)
func RoleMiddleware(db *pgxpool.Pool, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get("userID")
		if !exists {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		userID := userIDVal.(uint64)

		// Fetch user roles
		query := `
			SELECT r.name
			FROM roles r
			INNER JOIN user_roles ur ON ur.role_id = r.id
			WHERE ur.user_id = $1
		`
		rows, err := db.Query(context.Background(), query, userID)
		if err != nil {
			response.Error(c, http.StatusForbidden, "Failed to fetch user roles", err)
			c.Abort()
			return
		}
		defer rows.Close()

		userRoles := []string{}
		for rows.Next() {
			var role string
			if err := rows.Scan(&role); err != nil {
				response.Error(c, http.StatusInternalServerError, "Error scanning roles", err)
				c.Abort()
				return
			}
			userRoles = append(userRoles, role)
		}

		// Check if any of the required roles exist in user's roles
		allowed := false
		for _, r := range requiredRoles {
			if slices.Contains(userRoles, r) {
				allowed = true
				break
			}
		}

		if !allowed {
			response.Error(c, http.StatusForbidden, "Insufficient permissions", nil)
			c.Abort()
			return
		}

		c.Next()
	}

}
