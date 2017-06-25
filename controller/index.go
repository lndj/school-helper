package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
)

//The index page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"ip": c.ClientIP(),
	})

}
