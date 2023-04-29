package controller

import (
	"encoding/json"
	"net/http"

	"stac/database"
	"stac/models"
	"stac/parser"
	"stac/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
	"google.golang.org/protobuf/proto"
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
		handlePushEvent(hook, c)
	default:
		c.JSON(http.StatusNotImplemented, gin.H{"msg": "Received not implemented event type"})
		return
	}
}

func handlePushEvent(hook *parser.Webhook, c *gin.Context) {
	evt := github.PushEvent{}
	err := json.Unmarshal(hook.Payload, &evt)
	if utils.CheckError(err) {
		return
	}
	// lookup if this repo is in the database
	registered, err := database.DB.Has([]byte(evt.GetRepo().GetFullName()), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, OPServerError)
		return
	}
	if registered {
		// check if it uses secret
		pb, err := database.DB.Get([]byte(evt.GetRepo().GetFullName()), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		p := &models.GithubHook{}
		if err := proto.Unmarshal(pb, p); err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		if p.UseSecret {
			if !hook.Verify([]byte(utils.Config.Pwd)) {
				c.JSON(http.StatusUnauthorized, OPUnauth)
			}
		}

		// TODO: Execute CI/CD logic

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Repo not registered with stac, please use the register API"})
	}
}
