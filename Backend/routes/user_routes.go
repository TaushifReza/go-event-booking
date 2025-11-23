package routes

import (
	"github.com/TaushifReza/go-event-booking-api/controllers"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(server *gin.Engine, db *gorm.DB){
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	auth := server.Group("/auth")
	auth.POST("/register/", userController.Register)
	auth.POST("/login/", userController.Login)
}