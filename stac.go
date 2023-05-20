package main

import (
	"embed"
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

//go:embed template/*
var templates embed.FS

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the config file path")
	}

	// Inject Frontend code
	content, err := templates.ReadFile("template/main.html")
	if utils.CheckError(err) {
		panic("Frontend read error")
	}
	controller.MainContent = content

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
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
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
	err = router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	controller.Register(router)

	wg.Wait()
	log.Fatal(router.Run(utils.Config.IP + ":" + utils.Config.Port))
}
