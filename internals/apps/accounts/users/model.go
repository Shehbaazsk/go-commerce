package users

import "time"

type CreateUserRequest struct {
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    *string    `json:"last_name,omitempty"`
	Email       string     `json:"email" binding:"required,email"`
	Password    string     `json:"password" binding:"required,min=8"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
}

type UpdateUserRequest struct {
	FirstName   *string    `json:"first_name"`
	Email       *string    `json:"email" binding:"omitempty,email"`
	LastName    *string    `json:"last_name"`
	PhoneNumber *string    `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	IsActive    *bool      `json:"is_active"`
}

type UserResponse struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    *string    `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	IsActive    *bool      `json:"is_active"`
}
