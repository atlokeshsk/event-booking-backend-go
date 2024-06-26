package routes

import (
	"github.com/atlokeshsk/event-booking/middlewars"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the server.
// It initializes the routes for events and users.
//
// Parameters:
//   - server: A pointer to the gin.Engine instance where the routes will be registered.
func RegisterRoutes(server *gin.Engine) {
	registerEventsRoutes(server)
	registerUsersRoutes(server)
}

// registerEventsRoutes sets up the routes for event-related endpoints.
// It defines both public and authenticated routes for handling events.
//
// Public Routes:
// - GET /events/:id: Retrieves an event by its ID.
// - GET /events: Retrieves a list of events.
//
// Authenticated Routes (require authentication):
// - POST /events: Creates a new event.
// - PUT /events/:id: Updates an event by its ID.
// - DELETE /events/:id: Deletes an event by its ID.
// - POST /events/:id/register: Registers for an event by its ID.
// - DELETE /events/:id/register: Cancels registration for an event by its ID.
//
// Parameters:
// - server (*gin.Engine): The Gin engine instance to which the routes will be registered.
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
