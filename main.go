package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"school-helper/wechat"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Any("/wechat", wechat.WechatHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "This is a Wechat Server, powered by Golang.")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
