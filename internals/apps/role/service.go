package role

import (
	"context"
	"slices"

	db "github.com/shehbaazsk/go-commerce/db/queries"
	"github.com/shehbaazsk/go-commerce/internals/common/converters"
	"github.com/shehbaazsk/go-commerce/internals/constants"
	"github.com/shehbaazsk/go-commerce/utils"
)

func mapRoleToResponse(role db.Role) RoleResponse {
	return RoleResponse{
		ID:          int64(role.ID),
		Name:        role.Name,
		Description: converters.StringOrNil(role.Description),
		CreatedAt:   converters.TimeOrNil(role.CreatedAt),
		UpdatedAt:   converters.TimeOrNil(role.UpdatedAt),
	}
}

type Service interface {
	Create(ctx context.Context, req RoleRequest) (RoleResponse, error)
	Update(ctx context.Context, id int64, req UpdateRoleRequest) (RoleResponse, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (RoleResponse, error)
	GetAll(ctx context.Context, userId int64) ([]RoleResponse, error)
}

type service struct {
	queries *db.Queries
}

func NewRoleService(q *db.Queries) Service {
	return &service{queries: q}
}

func (s *service) Create(ctx context.Context, req RoleRequest) (RoleResponse, error) {
	roleParams := db.CreateRoleParams{
		Name:        req.Name,
		Description: converters.StringToNullString(req.Description),
	}
	role, err := s.queries.CreateRole(ctx, roleParams)
	if err != nil {
		return RoleResponse{}, err
	}
	return mapRoleToResponse(role), nil
}
func (s *service) Update(ctx context.Context, id int64, req UpdateRoleRequest) (RoleResponse, error) {
	roleParams := db.UpdateRoleParams{
		ID:          int32(id),
		Name:        req.Name,
		Description: converters.StringToNullString(req.Description),
		IsActive:    converters.BoolToNullBool(req.IsActive),
	}
	role, err := s.queries.UpdateRole(ctx, roleParams)
	if err != nil {
		return RoleResponse{}, err
	}
	return mapRoleToResponse(role), nil
}
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.queries.DeleteRole(ctx, int32(id))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetByID(ctx context.Context, id int64) (RoleResponse, error) {
	role, err := s.queries.GetRoleByID(ctx, int32(id))
	if err != nil {
		return RoleResponse{}, err
	}
	return mapRoleToResponse(role), nil
}

func (s *service) GetAll(ctx context.Context, userID int64) ([]RoleResponse, error) {
	userRoles, err := utils.GetUserRoles(ctx, s.queries, userID)
	if err != nil {
		return nil, err
	}
	var roles []db.Role
	if slices.Contains(userRoles, constants.RoleAdmin) {
		roles, err = s.queries.ListRoles(ctx)
	} else if slices.Contains(userRoles, constants.RoleStaff) {
		roles, err = s.queries.ListRolesWithoutAdminAndStaff(ctx)
	} else {
		roles, err = s.queries.ListRolesWithoutAdmin(ctx)
	}
	if err != nil {
		return nil, err
	}
	var roleResponse []RoleResponse
	for _, role := range roles {
		roleResponse = append(roleResponse, mapRoleToResponse(role))
	}
	return roleResponse, nil
}
