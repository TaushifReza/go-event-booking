package utils

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors(server *gin.Engine) {
	frontend := os.Getenv("FRONTEND_URL")

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontend},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}