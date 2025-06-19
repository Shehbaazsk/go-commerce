package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/common/response"
)

type Handler struct {
	service Service
}

func NewHandler(dbPool *pgxpool.Pool) *Handler {
	return &Handler{service: NewCustomerService(dbPool)}
}

// Crate a new customer
func (h *Handler) CreateCustomer(c *gin.Context) {
	var createCustomerReq CreateCustomerRequest
	if err := c.ShouldBindJSON(&createCustomerReq); err != nil {
		response.ValidationError(c, err)
		return
	}
	customerRes, err := h.service.CreateCustomer(c.Request.Context(), createCustomerReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create customer", err)
	}
	response.Success(c, http.StatusCreated, "Customer created successfully", customerRes)
}

// update a customer
func (h *Handler) UpdateCustomer(c *gin.Context) {
	idParam := c.Param("user_id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing User ID", nil)
		return
	}
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid User ID", err)
		return
	}
	var updateCustomerReq UpdateCustomerRequest
	if err := c.ShouldBindJSON(&updateCustomerReq); err != nil {
		response.ValidationError(c, err)
		return
	}
	customerRes, err := h.service.UpdateCustomer(c.Request.Context(), userID, updateCustomerReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update customer", err)
	}
	response.Success(c, http.StatusAccepted, "Customer updated successfully", customerRes)
}

// delete a customer
func (h *Handler) DeleteCustomer(c *gin.Context) {
	idParam := c.Param("user_id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing User ID", nil)
		return
	}
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid User ID", err)
		return
	}
	if err := h.service.DeleteCustomer(c.Request.Context(), userID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete customer", err)
		return
	}
	response.Success(c, http.StatusNoContent, "Customer deleted successfully", nil)
}

// get customer by user id
func (h *Handler) GetCustomerByUserID(c *gin.Context) {
	idParam := c.Param("user_id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing User ID", nil)
		return
	}
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid User ID", err)
		return
	}
	customerRes, err := h.service.GetCustomerByUserId(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get customer", err)
	}
	response.Success(c, http.StatusOK, "Customer retrieved successfully", customerRes)
}

// list all customer
func (h *Handler) ListCustomers(c *gin.Context) {
	var listCustomerReq ListCustomerRequest
	if err := c.ShouldBindJSON(&listCustomerReq); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Customer Request", err)
	}
	customersRes, err := h.service.GetAllCustomer(c.Request.Context(), listCustomerReq)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get customers", err)
	}
	response.Success(c, http.StatusOK, "Customers retrieved successfully", customersRes)
}
