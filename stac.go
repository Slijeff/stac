package main

import (
	"stac/controller"
	"stac/utils"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the config file path")
	}
	config := utils.ReadConfig(os.Args[1])

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/hook", controller.HandleGithubWebhook)
	router.Run(config.IP + ":" + config.Port)
}
