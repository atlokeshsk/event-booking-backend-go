package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	registerEventsRoutes(server)
	registerUsersRoutes(server)
}
