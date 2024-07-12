package main

import (
	"gin-jwt/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New() //可自定义中间件
	r.Use(gin.Logger(), gin.Recovery())

	routers.InitRouter(r)

	r1 := r.Group("/v1")
	r2 := r.Group("/v2")
	r1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r2.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
