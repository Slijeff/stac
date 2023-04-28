package controller

import (
	"encoding/json"
	"fmt"

	"stac/parser"
	"stac/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
)

func HandleGithubWebhook(c *gin.Context) {
	// hook, err := parser.Parse([]byte("safsdfasdf"), c.Request)
	hook, err := parser.ParseWithoutSecret(c.Request)
	if utils.CheckError(err) {
		return
	}
	// Determine which kind of event it is
	switch hook.Event {
	case "push":
		handlePushEvent(hook)
	default:
		fmt.Println("Received not implemented event type: ", hook.Event)
	}

	c.JSON(200, gin.H{"copythat!": c.GetHeader("X-GitHub-Delivery")})
}

func handlePushEvent(hook *parser.Webhook) {
	evt := github.PushEvent{}
	err := json.Unmarshal(hook.Payload, &evt)
	if utils.CheckError(err) {
		return
	}
	fmt.Println("receidved", evt.GetHeadCommit().GetAuthor().GetName())
}
