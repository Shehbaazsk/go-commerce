package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generic Response Format
type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

// Success Response Helper
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

// Error Response Helper
func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	})
}

// Validation Error Response
func ValidationError(c *gin.Context, err interface{}) {
	c.JSON(http.StatusUnprocessableEntity, APIResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Validation failed",
		Error:      err,
	})
}

// Unauthorized Error Response
func UnauthorizedError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	})
}

// Forbidden Error Response
func ForbiddenError(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		StatusCode: http.StatusForbidden,
		Message:    message,
	})
}

// Not Found Error Response
func NotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	})
}

// Bad Request Error Response
func BadRequestError(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusBadRequest, APIResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Error:      err,
	})
}
