package main

import (
	"log"
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
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptChan
		// do cleanups
		err := database.DB.Close()
		if err != nil {
			return
		}
		os.Exit(0)
	}()

	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	controller.Register(router)

	log.Fatal(router.Run(utils.Config.IP + ":" + utils.Config.Port))
}
