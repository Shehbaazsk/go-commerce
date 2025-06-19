package users

import (
	"context"

	db "github.com/shehbaazsk/go-commerce/db/queries"
	"github.com/shehbaazsk/go-commerce/internals/common/converters"
)

func mapUserToResponse(user db.User) UserResponse {
	return UserResponse{
		ID:          int(user.ID),
		FirstName:   user.FirstName,
		LastName:    converters.TextOrNil(user.LastName),
		Email:       user.Email,
		PhoneNumber: converters.TextOrNil(user.PhoneNumber),
		DateOfBirth: converters.DateOrNil(user.DateOfBirth),
		IsActive:    converters.BoolOrNil(user.IsActive),
	}
}

type Service interface {
	CreateUser(ctx context.Context, tx *db.Queries, req db.CreateUserParams) (UserResponse, error)
	UpdateUser(ctx context.Context, tx *db.Queries, req db.UpdateUserParams) (UserResponse, error)
	DeleteUser(ctx context.Context, tx *db.Queries, id int) error
}

type service struct{}

func NewUserService() Service {
	return &service{}
}

func (s *service) CreateUser(ctx context.Context, q *db.Queries, req db.CreateUserParams) (UserResponse, error) {
	user, err := q.CreateUser(ctx, req)
	if err != nil {
		return UserResponse{}, err
	}
	return mapUserToResponse(user), nil
}

func (s *service) UpdateUser(ctx context.Context, q *db.Queries, req db.UpdateUserParams) (UserResponse, error) {
	user, err := q.UpdateUser(ctx, req)
	if err != nil {
		return UserResponse{}, err
	}
	return mapUserToResponse(user), nil
}

func (s *service) DeleteUser(ctx context.Context, q *db.Queries, id int) error {
	return q.DeleteUser(ctx, int32(id))
}
