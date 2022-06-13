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

	return router
}
