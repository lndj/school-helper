package router

import (
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/lndj/school-helper/alert"
	"github.com/lndj/school-helper/controller"
	"github.com/lndj/school-helper/wechat"
)

//Loads all the router
func Load(middleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(middleware...)

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	appRoot, _ := os.Getwd()
	favicon := filepath.Join(appRoot, "/assets/favicon.ico")
	r.StaticFile("/favicon.ico", favicon)

	r.Any("/wechat", wechat.WechatHandler)

	r.GET("/", controller.Index)
	r.POST("/", controller.IndexPost)

	r.POST("/slack", alert.InteractionHandler)

	return r
}
