package main

import (
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
	"school-helper/wechat"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
)

func main() {
	r := gin.Default()

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

	r.Run(":" + port()) // listen and serve on 0.0.0.0:8080
}

//获取端口号
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	return port
}
