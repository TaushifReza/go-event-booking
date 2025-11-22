package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TaushifReza/go-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func eventRegister(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
		return
	}
	userId := c.GetInt64("id")

	_, err = models.GetEvent(eventID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Event with ID %d not found", eventID)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = models.RegisterForEvent(eventID, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
}