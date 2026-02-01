package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/TaushifReza/go-event-booking-api/controllers"
	"github.com/TaushifReza/go-event-booking-api/middleware"
	"github.com/TaushifReza/go-event-booking-api/repositories"
	"github.com/TaushifReza/go-event-booking-api/services"
)

func EventRegistrationRoutes(server *gin.Engine, db *gorm.DB) {
	repo := repositories.EventRegistrationRepository{DB: db}
	service := services.EventRegistrationService{Repo: &repo}
	controller := controllers.EventRegistrationController{EventRegistrationService: &service}

	api := server.Group("/api")

	protected := api.Group("/events").Use(middleware.AuthMiddleware())
	{
		protected.POST("/:id/registers/", controller.Register)
	}
}
