package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/common/response"
	"github.com/shehbaazsk/go-commerce/internals/constants"
)

// RBAC middleware (role-based access control)
func RoleMiddleware(db *pgxpool.Pool, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get(string(constants.UserIDKey))
		if !exists {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		userID := userIDVal.(int)

		userRoles, err := fetchUserRoles(db, userID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to fetch roles", err)
			c.Abort()
			return
		}

		if hasRequiredRole(userRoles, requiredRoles) {
			c.Next()
			return
		}

		response.Error(c, http.StatusForbidden, "Insufficient permissions", nil)
		c.Abort()
	}

}

func fetchUserRoles(db *pgxpool.Pool, userID int) ([]string, error) {
	query := `
        SELECT r.name
        FROM roles r
        INNER JOIN user_roles ur ON ur.role_id = r.id
        WHERE ur.user_id = $1
    `
	rows, err := db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []string{}
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func hasRequiredRole(userRoles, requiredRoles []string) bool {
	for _, required := range requiredRoles {
		for _, r := range userRoles {
			if r == required {
				return true
			}
		}
	}
	return false
}

// checkResourceOwnership verifies if the user owns the given resource.
func checkResourceOwnership(db *pgxpool.Pool, tableName string, userID int, resourceID int) (bool, error) {
	if tableName == "" {
		return false, errors.New("invalid resource table name")
	}

	query := `SELECT COUNT(1) FROM ` + tableName + ` WHERE id = $1 AND created_by = $2`
	var count int
	err := db.QueryRow(context.Background(), query, resourceID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func OwnerOnlyMiddleware(db *pgxpool.Pool, tableName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get(string(constants.UserIDKey))
		if !exists {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		userID := userIDVal.(int)

		resourceIDStr := c.Param("id")
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid resource ID", err)
			c.Abort()
			return
		}

		isOwner, err := checkResourceOwnership(db, tableName, userID, resourceID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed ownership check", err)
			c.Abort()
			return
		}
		if !isOwner {
			response.Error(c, http.StatusForbidden, "You don't own this resource", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

// RoleOrOwnerMiddleware checks if the user has one of the required roles
// or is the owner of the resource.
func RoleOrOwnerMiddleware(db *pgxpool.Pool, resourceTable string, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check for role
		roleMIddleware := RoleMiddleware(db, requiredRoles...)
		roleMIddleware(c)

		if c.IsAborted() {
			return
		}

		// check for ownership
		ownerMiddleware := OwnerOnlyMiddleware(db, resourceTable)
		ownerMiddleware(c)
	}
}
