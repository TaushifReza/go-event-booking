package middleware

import (
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        claims, err := utils.VerifyToken(token, "access")
        if err != nil {
            c.JSON(401, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.ID)
        c.Set("email", claims.Email)

        c.Next()
    }
}
