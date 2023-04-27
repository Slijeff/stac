package controller

import (
	"fmt"

	"stac/parser"

	"github.com/gin-gonic/gin"
)

func HandleGithubWebhook(c *gin.Context) {
	hook, err := parser.Parse([]byte("secret!"), c.Request)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(hook)
	c.Header("test", "success!")
	c.JSON(200, gin.H{"copythat!": c.GetHeader("token")})
}
