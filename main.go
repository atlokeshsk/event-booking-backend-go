package main

import (
	"log"

	"github.com/atlokeshsk/event-booking/db"
	"github.com/atlokeshsk/event-booking/routes"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application. It initializes the database,
// sets up the HTTP server with routes, and starts the server on port 8080.
func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
