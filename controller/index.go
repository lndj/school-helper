package controller

import (
	"github.com/lndj/school-helper/alert"
	"gopkg.in/gin-gonic/gin.v1"
)

// If use gin v2
//type FeedbackForm struct {
//	Ip      string `form:"ip"`
//	Sex     string `form:"sex"`
//	Message string `form:"message"`
//}

//The index page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"ip": c.ClientIP(),
	})
}

func AddFeedback(c *gin.Context) {
	msg := c.PostForm("message")
	alert.SendTextToSchoolHelperBot(msg)
	c.JSON(200, gin.H{"msg": msg})
}
