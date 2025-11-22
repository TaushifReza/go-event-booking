package main

import (
	"github.com/TaushifReza/go-event-booking-api/db"
	"github.com/TaushifReza/go-event-booking-api/routes"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
        panic("No .env file found (using environment variables)")
    }
	
	db.InitDB()
	server := gin.Default()

	utils.SetupCors(server)
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
