package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/constants"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

func RegisterPublicRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
	h := NewHandler(dbPool)
	customer := rg.Group("/customer")
	{
		customer.POST("/", h.CreateCustomer)
	}
}

func RegisterProtectedRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
	h := NewHandler(dbPool)
	customer := rg.Group("/customer")
	{

		customer.GET("/", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin, constants.RoleStaff), h.ListCustomers)
		customer.GET("/:id", middlewares.RoleOrOwnerMiddleware(dbPool, constants.UsersTable, constants.RoleAdmin, constants.RoleStaff), h.GetCustomerByUserID)
		customer.PATCH("/:id", middlewares.RoleOrOwnerMiddleware(dbPool, constants.UsersTable, constants.RoleAdmin), h.UpdateCustomer)
		customer.DELETE("/:id", middlewares.RoleOrOwnerMiddleware(dbPool, constants.UsersTable, constants.RoleAdmin), h.DeleteCustomer)

	}
}
