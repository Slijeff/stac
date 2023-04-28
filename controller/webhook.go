package controller

import (
	"encoding/json"
	"fmt"

	"stac/parser"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
)

func HandleGithubWebhook(c *gin.Context) {
	// hook, err := parser.Parse([]byte("safsdfasdf"), c.Request)
	hook, err := parser.ParseWithoutSecret(c.Request)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	evt := github.PushEvent{}
	err = json.Unmarshal(hook.Payload, &evt)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(evt.GetRepo().GetCloneURL())
	c.JSON(200, gin.H{"copythat!": c.GetHeader("token")})
}
