package main

import (
	"os/signal"
	"stac/controller"
	"stac/database"
	"stac/utils"
	"syscall"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the config file path")
	}
	config := utils.ReadConfig(os.Args[1])

	database.InitDB("./data")

	// gracefully exit
	interrupt_chan := make(chan os.Signal, 1)
	signal.Notify(interrupt_chan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt_chan
		// do cleanups
		database.DB.Close()
		os.Exit(0)
	}()
	database.DB.Put([]byte("key"), []byte("test-value"), nil)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/hook", controller.HandleGithubWebhook)
	router.Run(config.IP + ":" + config.Port)
}
