package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/apps/accounts/customer"
)

func RegisterPublicRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
	account := rg.Group("/account")
	{
		customer.RegisterPublicRoutes(account, dbPool)
	}
}

func RegisterProtectedRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
	account := rg.Group("/account")
	{
		customer.RegisterProtectedRoutes(account, dbPool)
	}
}
