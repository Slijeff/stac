package controller

import (
	"net/http"
	"stac/database"
	"stac/models"
	"stac/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

type RegisterRequestBody struct {
	// repo name in the format of User/Repo, or it can be an organization
	Name string `json:"name"`
	// whether to use secret for this repo or not
	Use_secret bool `json:"use_secret"`
}

func RegisterRepo(c *gin.Context) {

	if !verifyHeader(c) {
		return
	}

	var requestBody RegisterRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "request format is incorrect",
		})
		return
	}

	hasRepo, err := database.DB.Has([]byte(requestBody.Name), nil)
	if utils.CheckError(err) {
		c.JSON(http.StatusInternalServerError, OPServerError)
		return
	}
	if hasRepo {
		c.JSON(http.StatusOK, gin.H{
			"msg": "repo already exist",
		})
		return
	} else {
		// Create protobuf
		p := models.GithubHook{
			UseSecret: requestBody.Use_secret,
		}
		out, err := proto.Marshal(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		if err := database.DB.Put([]byte(requestBody.Name), out, nil); err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
	}
	c.JSON(http.StatusOK, OPSuccess)
}

func ChangeUseSecret(c *gin.Context) {
	if !verifyHeader(c) {
		return
	}

	var requestBody RegisterRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "request format is incorrect",
		})
		return
	}

	hasRepo, err := database.DB.Has([]byte(requestBody.Name), nil)
	if utils.CheckError(err) {
		c.JSON(http.StatusInternalServerError, OPServerError)
		return
	}
	if hasRepo {
		// Create protobuf
		p := models.GithubHook{
			UseSecret: requestBody.Use_secret,
		}
		out, err := proto.Marshal(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		if err := database.DB.Put([]byte(requestBody.Name), out, nil); err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		c.JSON(http.StatusOK, OPSuccess)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "repo don't exist, please register first",
		})
		return
	}

}

func verifyHeader(c *gin.Context) bool {
	stac_pwd := c.GetHeader("stac-pwd")
	if len(stac_pwd) == 0 {
		c.JSON(http.StatusUnauthorized, OPUnauth)
		return false
	}

	// might use secure compare?
	if stac_pwd != utils.Config.Pwd {
		c.JSON(http.StatusUnauthorized, OPUnauth)
		return false
	}

	return true
}
