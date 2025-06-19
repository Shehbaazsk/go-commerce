package role

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shehbaazsk/go-commerce/internals/constants"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

// func RegisterPublicRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
// 	// h := NewHandler(queries)
// 	roles := rg.Group("/roles")
// 	{
// 		roles.GET("/", func(c *gin.Context) {
// 			// implement logic to get all roles
// 		})
// 	}
// 	// other public routes for role module
// }

func RegisterProtectedRoutes(rg *gin.RouterGroup, dbPool *pgxpool.Pool) {
	h := NewHandler(dbPool)
	roles := rg.Group("/roles")
	{
		roles.POST("/", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin), h.CreateRole)
		roles.GET("/", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin, constants.RoleStaff), h.GetAllRoles)
		roles.GET("/:id", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin, constants.RoleStaff), h.GetRoleByID)
		roles.PATCH("/:id", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin), h.UpdateRole)
		roles.DELETE("/:id", middlewares.RoleMiddleware(dbPool, constants.RoleAdmin), h.DeleteRole)
	}
}
