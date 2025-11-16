package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!", "events": events, "count": len(events)})
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(event)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event.UserID = 1
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with ID %d not found", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello World!", "events": event})
}