package routes

import (
	"net/http"

	"github.com/atlokeshsk/event-booking/models"
	"github.com/atlokeshsk/event-booking/utils"
	"github.com/gin-gonic/gin"
)

// signup handles the user signup process. It binds the JSON payload to a User model,
// validates it, and attempts to save the user to the database. Appropriate HTTP responses
// are returned based on the success or failure of these operations.
func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messsage": http.StatusText(http.StatusBadRequest)})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messsage": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "user_id": user.ID})
}

// login handles the user login process. It binds the JSON payload to a user model,
// validates the credentials, generates a JWT token upon successful validation,
// and returns appropriate HTTP responses based on the outcome.
func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messsage": http.StatusText(http.StatusBadRequest)})
		return
	}

	err = user.ValidateCredential()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	token, err := utils.GenerateJwtToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}
