package routes

import (
	"go-line-demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetMessageRoutes(router *gin.Engine) *gin.Engine {
	message := router.Group("/messages")
	{
		// Receive messages
		message.POST("/receive", controllers.ReceiveMessage)
		// Send message
		message.POST("/push", controllers.PushMessage)
	}
	return router
}
