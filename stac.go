package main

import (
	"embed"
	"log"
	"os/signal"
	"stac/controller"
	"stac/database"
	"stac/utils"
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
	utils.ReadConfig(os.Args[1])
	database.InitDB("./data")

	// Setup Logging
	staclogger := utils.GetStacLogger()
	staclogger.Println("Starting...")

	// Inject Frontend code
	content, err := templates.ReadFile("template/main.html")
	if utils.CheckError(err) {
		panic("Frontend read error")
	}
	controller.MainContent = content

	// gracefully exit
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptChan
		staclogger.Println("Exiting...")
		// do cleanups
		err := database.DB.Close()
		if err != nil {
			return
		}
		utils.CloseAllLogFiles()
		os.Exit(0)
	}()

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	controller.Register(router)
	log.Fatal(router.Run(utils.Config.IP + ":" + utils.Config.Port))
}
