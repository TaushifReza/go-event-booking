package main

import (
	"github.com/TaushifReza/go-event-booking-api/db"
	"github.com/TaushifReza/go-event-booking-api/logger"
	"github.com/TaushifReza/go-event-booking-api/routes"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	
	if err := godotenv.Load(".env"); err != nil{
        panic("No .env file found (using environment variables)")
    }
	
	dbInstance, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	logger.Info("Successfully connect to database.")

	err = db.Migrate()

	logger.Info("Successfully migrate table to database.")

	if err != nil {
		panic(err)
	}

	server := gin.Default()

	utils.SetupCors(server)

	// add routes
	routes.UserRoutes(server, dbInstance)

	logger.Info("Server running on port 8080")
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
