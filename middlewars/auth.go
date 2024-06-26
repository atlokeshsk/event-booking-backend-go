package middlewars

import (
	"fmt"
	"net/http"

	"github.com/atlokeshsk/event-booking/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware function for the Gin framework that checks for an
// Authorization token in the request header. If the token is missing or invalid,
// it aborts the request with a 401 Unauthorized status. If the token is valid,
// it sets the user ID in the context and proceeds to the next handler.
//
// Parameters:
//
//	c (*gin.Context): The context for the current request, which includes the
//	                  request and response objects, as well as other metadata.
func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	fmt.Println(token)
	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	c.Set("user_id", userID)
	c.Next()
}
