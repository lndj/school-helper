package router

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lndj/school-helper/alert"
	"github.com/lndj/school-helper/config"
	"github.com/lndj/school-helper/wechat"
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
	appRoot, _ := os.Getwd()
	favicon := filepath.Join(appRoot, "/assets/favicon.ico")
	r.StaticFile("/favicon.ico", favicon)

	r.Any("/wechat", wechat.WechatHandler)

	r.GET("/", func(c *gin.Context) {
		fmt.Println(config.Configure.String("redis.addr"))
		c.String(200, "This is a Wechat Server, powered by Golang.")
	})

	r.POST("/slack", alert.InteractionHandler)

	return r
}
