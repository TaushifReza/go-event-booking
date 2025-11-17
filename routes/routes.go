package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/event/:id/", getEvent)
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvents)
}