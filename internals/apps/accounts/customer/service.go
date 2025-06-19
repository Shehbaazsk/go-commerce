package customer

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shehbaazsk/go-commerce/db/queries"
	"github.com/shehbaazsk/go-commerce/internals/apps/accounts/users"
	"github.com/shehbaazsk/go-commerce/internals/common/converters"
	"github.com/shehbaazsk/go-commerce/internals/constants"
	"github.com/shehbaazsk/go-commerce/utils"
)

func mapCustomerToResponse(user users.UserResponse, customer db.CustomerProfile) CustomerResponse {
	contactPreference, err := converters.FromJSONB(customer.ContactPreference)
	if err != nil {
		return CustomerResponse{}
	}
	return CustomerResponse{
		UserID:            user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		DateOfBirth:       user.DateOfBirth,
		ContactPreference: &contactPreference,
	}
}

type Service interface {
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CustomerResponse, error)
	UpdateCustomer(ctx context.Context, userID int, req UpdateCustomerRequest) (CustomerResponse, error)
	DeleteCustomer(ctx context.Context, userID int) error
	GetCustomerByUserId(ctx context.Context, userID int) (CustomerResponse, error)
	GetAllCustomer(ctx context.Context, req ListCustomerRequest) ([]CustomerResponse, error)
}

type service struct {
	dbPool  *pgxpool.Pool
	account users.Service
}

func NewCustomerService(dbPool *pgxpool.Pool) Service {
	return &service{dbPool: dbPool}
}

func (s *service) CreateCustomer(ctx context.Context, req CreateCustomerRequest) (CustomerResponse, error) {
	hashPasswd, err := utils.HashPassword(req.Password)
	if err != nil {
		return CustomerResponse{}, err
	}
	createUserParams := db.CreateUserParams{
		FirstName:    req.FirstName,
		LastName:     converters.StringToPgText(req.LastName),
		Email:        req.Email,
		PasswordHash: hashPasswd,
		PhoneNumber:  converters.StringToPgText(req.PhoneNumber),
	}

	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return CustomerResponse{}, err
	}
	qtx := db.New(tx)
	userRes, err := s.account.CreateUser(ctx, qtx, createUserParams)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	updateUserAuditParams := db.UpdateUserAuditFieldsParams{
		CreatedBy: converters.IntToPgInt4(&userRes.ID),
		UpdatedBy: converters.IntToPgInt4(&userRes.ID),
		ID:        int32(userRes.ID),
	}
	err = qtx.UpdateUserAuditFields(ctx, updateUserAuditParams)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	customerReq := db.CreateCustomerProfileParams{
		UserID:  int32(userRes.ID),
		Column2: req.ContactPreference,
	}
	customerRes, err := qtx.CreateCustomerProfile(ctx, customerReq)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	return mapCustomerToResponse(userRes, customerRes), nil
}

func (s *service) UpdateCustomer(ctx context.Context, userID int, req UpdateCustomerRequest) (CustomerResponse, error) {
	currentUserId := ctx.Value(constants.UserIDKey).(int)
	updateUserParams := db.UpdateUserParams{
		Email:       converters.StringToPgText(req.Email),
		FirstName:   converters.StringToPgText(req.FirstName),
		LastName:    converters.StringToPgText(req.LastName),
		PhoneNumber: converters.StringToPgText(req.PhoneNumber),
		DateOfBirth: converters.TimeToPgDate(req.DateOfBirth),
		IsActive:    converters.BoolToPgBool(req.IsActive),
		UpdatedBy:   converters.IntToPgInt4(&currentUserId),
		ID:          int32(userID),
	}
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return CustomerResponse{}, err
	}
	qtx := db.New(tx)
	userRes, err := s.account.UpdateUser(ctx, qtx, updateUserParams)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	contactPreference, err := converters.ToJSONB(*req.ContactPreference)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	updateCustomerReq := db.UpdateCustomerProfileParams{
		UserID:            int32(userID),
		ContactPreference: contactPreference,
	}
	customerRes, err := qtx.UpdateCustomerProfile(ctx, updateCustomerReq)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return CustomerResponse{}, err
	}
	return mapCustomerToResponse(userRes, customerRes), nil

}

func (s *service) DeleteCustomer(ctx context.Context, userID int) error {
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	qtx := db.New(tx)
	err = s.account.DeleteUser(ctx, qtx, userID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = qtx.DeleteCustomerProfile(ctx, int32(userID))
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return nil
}

func (s *service) GetCustomerByUserId(ctx context.Context, userID int) (CustomerResponse, error) {
	qtx := db.New(s.dbPool)
	currentUserId := ctx.Value(string(constants.UserIDKey)).(int)
	ownerID, err := qtx.GetUserIDByCreatedBy(ctx, converters.IntToPgInt4(&currentUserId))
	if err != nil {
		return CustomerResponse{}, err
	}
	if ownerID != int32(userID) {
		customerRes, err := qtx.GetCustomerByUserIDLimited(ctx, int32(userID))
		if err != nil {
			return CustomerResponse{}, err
		}
		return CustomerResponse{
			FirstName: customerRes.FirstName,
			LastName:  converters.TextOrNil(customerRes.LastName),
			Email:     customerRes.Email,
		}, nil
	}
	customerRes, err := qtx.GetCustomerByUserID(ctx, int32(userID))
	if err != nil {
		return CustomerResponse{}, err
	}
	contactPrefernece, err := converters.FromJSONB(customerRes.ContactPreference)
	if err != nil {
		return CustomerResponse{}, err
	}
	return CustomerResponse{
		UserID:            int(customerRes.UserID),
		FirstName:         customerRes.FirstName,
		LastName:          converters.TextOrNil(customerRes.LastName),
		Email:             customerRes.Email,
		PhoneNumber:       converters.TextOrNil(customerRes.PhoneNumber),
		DateOfBirth:       converters.DateOrNil(customerRes.DateOfBirth),
		ContactPreference: &contactPrefernece,
	}, nil
}

func (s *service) GetAllCustomer(ctx context.Context, req ListCustomerRequest) ([]CustomerResponse, error) {
	qtx := db.New(s.dbPool)
	req.SetDefaults()
	listCustomersParams := db.ListCustomersPaginatedParams{
		Limit:  int32(req.PerPage),
		Offset: int32(req.Page),
	}
	customersRes, err := qtx.ListCustomersPaginated(ctx, listCustomersParams)
	if err != nil {
		return nil, err
	}
	var customers []CustomerResponse
	for _, customer := range customersRes {
		contactPrefernece, err := converters.FromJSONB(customer.ContactPreference)
		if err != nil {
			return nil, err
		}
		customers = append(customers, CustomerResponse{
			UserID:            int(customer.UserID),
			FirstName:         customer.FirstName,
			LastName:          converters.TextOrNil(customer.LastName),
			Email:             customer.Email,
			PhoneNumber:       converters.TextOrNil(customer.PhoneNumber),
			DateOfBirth:       converters.DateOrNil(customer.DateOfBirth),
			ContactPreference: &contactPrefernece,
		})
	}
	return customers, nil
}
