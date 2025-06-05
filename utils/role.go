package utils

import (
	"context"

	db "github.com/shehbaazsk/go-commerce/db/queries"
)

func GetUserRoles(ctx context.Context, db *db.Queries, UserId int64) ([]string, error) {
	roles, err := db.GetUserRoles(ctx, int32(UserId))
	if err != nil {
		return nil, err
	}
	return roles, nil
}
