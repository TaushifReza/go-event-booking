package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func updateEvents(c *gin.Context){
	id , err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
	}

	_, err = models.GetEvent(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with ID %d not found", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBind(&updatedEvent)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, e := range errs {
				out[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.UpdateEvent()
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvents(c *gin.Context){
	id , err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
	}

	event , err := models.GetEvent(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with ID %d not found", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = event.DeleteEvent()
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}