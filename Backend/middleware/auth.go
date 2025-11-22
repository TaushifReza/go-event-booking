package middleware

import (
	"fmt"
	"net/http"

	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context){
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token required"})
		return
	}

	id, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println("ERROR: ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token."})
		return
	}

	c.Set("id", id)
	c.Next()
}