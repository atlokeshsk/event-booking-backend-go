package routes

import (
	"net/http"
	"strconv"

	"github.com/atlokeshsk/event-booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch the events right now"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": events})
}

func getEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request id should be a valid integer"})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})

}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not process the response body"})
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	userID := c.GetInt64("user_id")
	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "un authorized"})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
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
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
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
