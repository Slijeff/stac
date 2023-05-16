package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var MainContent []byte

func HandleMainFrontend(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", MainContent)
}
