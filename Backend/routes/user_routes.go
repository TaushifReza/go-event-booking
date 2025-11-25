package routes

import (
	"github.com/TaushifReza/go-event-booking-api/controllers"
	"github.com/TaushifReza/go-event-booking-api/middleware"
	"github.com/TaushifReza/go-event-booking-api/repositories"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(server *gin.Engine, db *gorm.DB){
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)
	userController := controllers.NewUserController(userService)

	auth := server.Group("/auth")
	{
		auth.POST("/register/", userController.Register)
		auth.POST("/login/", userController.Login)
	}

	api := server.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/user/me/", userController.GetUserInfo)
	}
}