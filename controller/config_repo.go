package controller

import (
	"fmt"
	"net/http"
	"stac/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	// repo name in the format of User/Repo, or it can be an organization
	Name string `json:"name"`
	// whether to use secret for this repo or not
	Use_secret bool `json:"use_secret"`
}

func RegisterRepo(c *gin.Context) {
	stac_pwd := c.GetHeader("stac-pwd")
	if len(stac_pwd) == 0 {
		fmt.Println("Please provide your stac_pwd when configuring repos")
		return
	}

	// might use secure compare?
	if stac_pwd != utils.Config.Pwd {
		c.JSON(http.StatusForbidden, gin.H{
			"code": http.StatusForbidden,
			"msg":  "password doesn't match",
		})
		return
	}
	var requestBody RegisterRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println("The request format is incorrect")
	}
	fmt.Println(requestBody)
	fmt.Println(requestBody.Name, requestBody.Use_secret)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}
