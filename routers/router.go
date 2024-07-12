package routers

import (
	"gin-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middleware.Cors(), middleware.Auth())
	//课程相关接口
	initCourse(api)
	//用户相关接口
	initUser(api)

	notAuthApi := router.Group("/api")
	notAuthApi.Use(middleware.Cors())
	//登录
	initLogin(notAuthApi)
}
