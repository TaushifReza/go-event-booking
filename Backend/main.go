package main

import (
	"fmt"

	"github.com/TaushifReza/go-event-booking-api/db"
	"github.com/TaushifReza/go-event-booking-api/routes"
	"github.com/TaushifReza/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil{
        panic("No .env file found (using environment variables)")
    }
	
	dbInstance, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connect to database.")

	err = db.Migrate()

	fmt.Println("Successfully migrate table to database.")

	if err != nil {
		panic(err)
	}

	server := gin.Default()

	utils.SetupCors(server)

	// add routes
	routes.UserRoutes(server, dbInstance)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
