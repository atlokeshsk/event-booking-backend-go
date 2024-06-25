package routes

import (
	"net/http"

	"github.com/atlokeshsk/event-booking/models"
	"github.com/atlokeshsk/event-booking/utils"
	"github.com/gin-gonic/gin"
)


func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request data"})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messsage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "user_id": user.ID})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to process the data bad request"})
		return
	}

	err = user.ValidateCredential()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateJwtToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}
