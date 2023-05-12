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
	UseSecret bool `json:"use_secret"`
}

type SetSecretBody struct {
	// repo name in the format of User/Repo, or it can be an organization
	Name string `json:"name"`
	// whether to use secret for this repo or not
	Secret string `json:"secret"`
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
			UseSecret: requestBody.UseSecret,
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
		// get protobuf
		pb := &models.GithubHook{}
		if err := utils.GetProtoFromDB(requestBody.Name, pb); err != nil {
			c.JSON(http.StatusInternalServerError, OPServerError)
			return
		}
		pb.UseSecret = requestBody.UseSecret
		// store it back
		if err := utils.PutProtoToDB(requestBody.Name, pb); err != nil {
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
	stacPwd := c.GetHeader("stac-pwd")
	if len(stacPwd) == 0 {
		c.JSON(http.StatusUnauthorized, OPUnauth)
		return false
	}

	// might use secure compare?
	if stacPwd != utils.Config.Pwd {
		c.JSON(http.StatusUnauthorized, OPUnauth)
		return false
	}

	return true
}

func SetSecret(c *gin.Context) {
	if !verifyHeader(c) {
		return
	}

	var requestBody SetSecretBody

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
	if !hasRepo {
		c.JSON(http.StatusOK, gin.H{
			"msg": "repo don't exist, please register first",
		})
		return
	}
	// get protobuf
	pb := &models.GithubHook{}
	if err := utils.GetProtoFromDB(requestBody.Name, pb); err != nil {
		c.JSON(http.StatusInternalServerError, OPServerError)
		return
	}
	pb.Secret = requestBody.Secret
	// store it back
	if err := utils.PutProtoToDB(requestBody.Name, pb); err != nil {
		c.JSON(http.StatusInternalServerError, OPServerError)
		return
	}
	c.JSON(http.StatusOK, OPSuccess)
}
