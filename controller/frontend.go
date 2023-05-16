package controller

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Content string

func HandleMainFrontend(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(Content))
}
