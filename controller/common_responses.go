package controller

import "github.com/gin-gonic/gin"

var (
	OPSuccess     = gin.H{"msg": "success"}
	OPServerError = gin.H{"msg": "server error"}
	OPUnauth      = gin.H{"msg": "unauthorized operation"}
)

func OPCustomErr(e error) gin.H {
	return gin.H{"msg": e.Error()}
}
