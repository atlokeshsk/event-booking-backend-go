package routes

import (
	"net/http"
	"strconv"

	"github.com/atlokeshsk/event-booking/models"
	"github.com/gin-gonic/gin"
)

// getEvents handles the HTTP request to retrieve all events.
// @Summary Get all events
// @Description Fetches all events from the database and returns them in the response.
// @Tags events
// @Produce json
// @Success 200 {object} gin.H{"message": []models.Event}
// @Failure 500 {object} gin.H{"message": string}
// @Router /events [get]
func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": events})
}

// getEventByID handles the HTTP request to retrieve an event by its ID.
// @param c *gin.Context - the context of the HTTP request, which includes parameters and other metadata.
// @return JSON response with the event data if found, or an error message if not found or if the ID is invalid.
func getEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
}

// createEvent handles the creation of a new event.
// It binds the JSON payload to an Event model, sets the user ID from the context,
// and attempts to save the event to the database. Appropriate JSON responses
// are returned based on the success or failure of these operations.
func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}
	userID := c.GetInt64("user_id")
	event.UserID = userID
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create the event at the moment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event_id": event.ID})
}

func updateEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not present"})
		return
	}
	userID := c.GetInt64("user_id")
	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "un authorized"})
		return
	}
	var updatedEvent *models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})

}

func deleteEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not present for the id"})
		return
	}

	userID := c.GetInt64("user_id")
	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "un authorized"})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event deleted from the db"})
}

func registerForEvent(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not present"})
		return
	}
	userID := c.GetInt64("user_id")
	err = event.Register(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to register for the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered for the event successfully"})
}

func cancelRegistration(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not present"})
		return
	}
	userID := c.GetInt64("user_id")
	err = event.CancelRegistration(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to cancel the registration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration cancelled succesffullty"})
}
