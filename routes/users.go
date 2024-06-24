package routes

import (
	"net/http"

	"github.com/atlokeshsk/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerUsersRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
}

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
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message": "unable to process the data bad request"})
		return
	}

	
	
}
