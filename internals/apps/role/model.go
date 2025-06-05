package role

import "time"

type RoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
}

type RoleResponse struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type UpdateRoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}
