package controller

import (
	"github.com/gin-gonic/gin"
)

func HandleGithubWebhook(c *gin.Context) {
	c.Header("test", "success!")
	c.JSON(200, gin.H{"copythat!": c.GetHeader("token")})
}
