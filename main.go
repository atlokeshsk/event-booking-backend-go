package main

import (
	"github.com/atlokeshsk/event-booking/db"
	"github.com/atlokeshsk/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
