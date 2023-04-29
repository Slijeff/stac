package controller

import "github.com/gin-gonic/gin"

var (
	OPSuccess     = gin.H{"msg": "success"}
	OPServerError = gin.H{"msg": "server error"}
	OPUnauth      = gin.H{"msg": "unauthorized operation"}
)
