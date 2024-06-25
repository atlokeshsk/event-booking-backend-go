package routes

import (
	"github.com/atlokeshsk/event-booking/middlewars"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	registerEventsRoutes(server)
	registerUsersRoutes(server)
}

func registerEventsRoutes(server *gin.Engine) {
	server.GET("/events/:id", getEventByID)
	server.GET("/events", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewars.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEventByID)
	authenticated.DELETE("/events/:id", deleteEventByID)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}

func registerUsersRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
}
