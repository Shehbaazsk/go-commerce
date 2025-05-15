package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/controllers"
	"github.com/shehbaazsk/go-commerce/services"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	users := rg.Group("/users")
	{
		users.POST("", userController.CreateUser)
		users.GET("", userController.GetAllUsers)
		users.GET("/:id", userController.GetUserByID)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
}
