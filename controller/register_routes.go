package controller

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// Register all routes here

	// WEBHOOK REQUEST GROUP
	hookRoute := router.Group("/hook")
	hookRoute.POST("/", HandleGithubWebhook)

	// CONFIG SETTING GROUP
	configRoute := router.Group("/config")
	configRoute.POST("/register", RegisterRepo)
	configRoute.POST("/secret", ChangeUseSecret)
	configRoute.POST("/setsec", SetSecret)

	// DATABASE OP GROUP
	dbRoute := router.Group("/db")
	dbRoute.GET("/getall", GetAllFromDB)
	dbRoute.GET("/delall", DeleteAllFromDB)
	dbRoute.GET("/delkey", DeleteSingleKey)

	// FRONTEND GROUP
	router.GET("/f", HandleMainFrontend)
}
