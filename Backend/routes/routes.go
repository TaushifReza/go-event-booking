package routes

import (
	"github.com/TaushifReza/go-event-booking-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)
	server.GET("/event/:id/", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvents)
	authenticated.DELETE("/events/:id", deleteEvents)

	server.POST("/signup", signup)
	server.POST("/login", login)
}