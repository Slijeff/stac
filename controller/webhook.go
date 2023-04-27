package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleGithubWebhook(c *gin.Context) {
	fmt.Printf("token: %s\n", c.GetHeader("token"))
	c.Header("test", "success!")
}
