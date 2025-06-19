package utils

import (
	"context"

	db "github.com/shehbaazsk/go-commerce/db/queries"
)

func GetUserRoles(ctx context.Context, db *db.Queries, UserId int) ([]string, error) {
	roles, err := db.GetUserRoles(ctx, int32(UserId))
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func HasRole(userRoles, requiredRoles []string) bool {
	for _, required := range requiredRoles {
		for _, r := range userRoles {
			if r == required {
				return true
			}
		}
	}
	return false
}
