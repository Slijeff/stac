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
	utils.ReadConfig(os.Args[1])

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

	router := gin.Default()
	router.SetTrustedProxies(nil)

	controller.Register(router)

	router.Run(utils.Config.IP + ":" + utils.Config.Port)
}
