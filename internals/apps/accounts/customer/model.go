package customer

import "time"

type CreateCustomerRequest struct {
	FirstName         string                  `json:"first_name" binding:"required"`
	LastName          *string                 `json:"last_name"`
	Email             string                  `json:"email" binding:"required,email"`
	Password          string                  `json:"password" binding:"required,min=8"`
	PhoneNumber       *string                 `json:"phone_number,omitempty"`
	DateOfBirth       *time.Time              `json:"date_of_birth,omitempty"`
	ContactPreference *map[string]interface{} `json:"contact_preference,omitempty"`
}

type UpdateCustomerRequest struct {
	FirstName         *string                 `json:"first_name"`
	LastName          *string                 `json:"last_name"`
	Email             *string                 `json:"email"`
	PhoneNumber       *string                 `json:"phone_number"`
	DateOfBirth       *time.Time              `json:"date_of_birth"`
	IsActive          *bool                   `json:"is_active"`
	ContactPreference *map[string]interface{} `json:"contact_preference"`
}

type CustomerResponse struct {
	UserID            int                     `json:"id"`
	FirstName         string                  `json:"first_name"`
	LastName          *string                 `json:"last_name"`
	Email             string                  `json:"email"`
	PhoneNumber       *string                 `json:"phone_number"`
	DateOfBirth       *time.Time              `json:"date_of_birth"`
	ContactPreference *map[string]interface{} `json:"contact_preference"`
}

type ListCustomerRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (r *ListCustomerRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PerPage <= 0 {
		r.PerPage = 10
	}
}
