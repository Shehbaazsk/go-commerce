package role

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
	return &Handler{service: NewRoleService(dbPool)}
}

// Create a new role
func (h *Handler) CreateRole(c *gin.Context) {
	var req RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	role, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create role", err)
		return
	}
	response.Success(c, http.StatusCreated, "Role created successfully", role)
}

func (h *Handler) UpdateRole(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing role ID", nil)
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 60)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	role, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update role", err)
		return
	}
	response.Success(c, http.StatusAccepted, "Role updated successfully", role)
}

func (h *Handler) DeleteRole(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing role ID", nil)
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 60)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete role", err)
	}
	response.Success(c, http.StatusNoContent, "Role deleted successfully", nil)
}

func (h *Handler) GetRoleByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		response.Error(c, http.StatusBadRequest, "Missing role ID", nil)
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 60)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	role, err := h.service.GetByID(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get role", err)
	}
	response.Success(c, http.StatusOK, "Role retrieved successfully", role)
}

func (h *Handler) GetAllRoles(c *gin.Context) {
	userIDVal, exists := c.Get("userIID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userID := userIDVal.(int64)
	roles, err := h.service.GetAll(c, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get roles", err)
	}
	response.Success(c, http.StatusOK, "Roles retrieved successfully", roles)

}
