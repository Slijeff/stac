package main

import (
	"fmt"
	"stac/controller"
	"stac/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	config := utils.ReadConfig("./config.json")
	fmt.Println("IP: ", config.IP, "Port: ", config.Port)

	router := gin.Default()
	router.GET("/hook", controller.HandleGithubWebhook)
	router.Run(config.IP + ":" + config.Port)
}
