package main

import (
	"fmt"
	"go-line-demo/config"
	"go-line-demo/database"
	"go-line-demo/routes"
)

func main() {
	config.Init()
	database.Init()
	router := routes.SetupRouter()
	config := config.GetConfig()

	router.Run(fmt.Sprintf(":%v", config.GetString("PORT")))
	defer database.Close()
}
