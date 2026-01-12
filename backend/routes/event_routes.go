package routes

import (
	"github.com/TaushifReza/go-event-booking-api/controllers"
	"github.com/TaushifReza/go-event-booking-api/middleware"
	"github.com/TaushifReza/go-event-booking-api/repositories"
	"github.com/TaushifReza/go-event-booking-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EventRoutes(server *gin.Engine, db *gorm.DB) {
	repo := repositories.NewEventRepository(db)
	eventService := services.NewEventService(repo)
	eventController := controllers.NewEventController(eventService)

	api := server.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/events/", eventController.Create)
		api.GET("/events/", eventController.GetAll)
	}
}
