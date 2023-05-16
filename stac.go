package main

import (
	_ "embed"
	"log"
	"os/signal"
	"stac/controller"
	"stac/database"
	"stac/utils"
	"sync"
	"syscall"

	"os"

	"github.com/gin-gonic/gin"
)

//go:embed template/main.html
var html string

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the config file path")
	}
	controller.Content = html
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		utils.ReadConfig(os.Args[1])
		wg.Done()
	}()
	go func() {
		database.InitDB("./data")
		wg.Done()
	}()

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

	wg.Wait()
	log.Fatal(router.Run(utils.Config.IP + ":" + utils.Config.Port))
}
