package routers

import (
	"gin-jwt/login"
	"github.com/gin-gonic/gin"
)

func initLogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/login", login.Login)
	}
}
