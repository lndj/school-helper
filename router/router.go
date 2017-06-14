package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"school-helper/wechat"
)

//Loads all the router
func Load(middleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(middleware...)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.Any("/wechat", wechat.WechatHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "This is a Wechat Server, powered by Golang.")
	})

	return r
}
