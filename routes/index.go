package routes

import (
	"go-line-demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Ping test
	router.GET("/", controllers.GetHi)

	// message routes
	router = SetMessageRoutes(router)

	// Get message list of user
	router.GET("/user/:userid/messages", controllers.GetMessages)

	return router
}
