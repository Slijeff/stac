package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"stac/database"
	"stac/models"
	"stac/parser"
	"stac/utils"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v52/github"
	"google.golang.org/protobuf/proto"
)

func HandleGithubWebhook(c *gin.Context) {
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
			// TODO: should get password from database
			if !hook.Verify([]byte(utils.Config.Pwd)) {
				c.JSON(http.StatusUnauthorized, OPUnauth)
				return
			}
		}

		// Execute CI/CD logic
		cloneURL := evt.GetRepo().GetCloneURL()
		repoPath := path.Join(utils.Config.Base, evt.GetRepo().GetName())
		existed, err := exists(repoPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		if existed {
			// try pulling the repo if it already exists
			localRepo, err := git.PlainOpen(repoPath)
			if utils.CheckError(err) {
				c.JSON(http.StatusInternalServerError, OPCustomErr(err))
				return
			}
			w, err := localRepo.Worktree()
			if utils.CheckError(err) {
				c.JSON(http.StatusInternalServerError, OPCustomErr(err))
				return
			}
			err = w.Pull(&git.PullOptions{
				Progress: os.Stdout,
			})
			if utils.CheckError(err) {
				// TODO: This is only used for debugging
				//c.JSON(http.StatusInternalServerError, OPCustomErr(err))
				//return
			}
		} else {
			_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
				URL:      cloneURL,
				Progress: os.Stdout,
			})
			if utils.CheckError(err) {
				c.JSON(http.StatusInternalServerError, OPCustomErr(err))
				return
			}
		}

		// parse Stacfile
		stages, err := parser.ParseYaml(path.Join(repoPath, "Stacfile.yaml"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		// start goroutines for each `routine` type commands
		var wg sync.WaitGroup
		for _, stage := range stages {
			if stage.Type == "routine" {
				wg.Add(1)
				// goroutines
				go func(s parser.Stage) {
					defer wg.Done()
					for _, c := range s.Commands {
						comms := strings.Fields(c)
						out, err := exec.Command("powershell", comms...).Output()
						if utils.CheckError(err) {
							return
						}
						fmt.Printf("%s\n", out)
					}
				}(stage)
			}
		}
		wg.Wait()

		// TODO: wait for concurrent execute, then start doing sequential

		c.JSON(http.StatusOK, OPSuccess)
		return

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Repo not registered with stac, please use the register API"})
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
