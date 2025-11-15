package main

import (
	"fmt"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context){
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!", "events": events, "count": len(events)})
}

func createEvent(c *gin.Context){
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil{
		fmt.Println(event)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}