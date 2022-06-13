package main

import (
	"fmt"
	"go-line-demo/config"
	"go-line-demo/database"
	"go-line-demo/routes"
	"go-line-demo/utils"
	"go-line-demo/validators"
)

func main() {
	config.Init()
	database.Init()
	utils.NewLinebot()
	validators.RegisterValidation()
	router := routes.SetupRouter()
	config := config.GetConfig()

	router.Run(fmt.Sprintf(":%v", config.GetString("PORT")))
	defer database.Close()
}
