package controller

import (
	"fmt"
	"github.com/lndj/school-helper/alert"
	"github.com/lndj/school-helper/utils"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
)

//The index page
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"ip": c.ClientIP(),
	})
}

//Handle the index page form submit
func IndexPost(c *gin.Context) {
	userIP := c.ClientIP()
	sex := c.PostForm("sex")
	content, ok := c.GetPostForm("content")

	if ok {

		sexInt, err := strconv.Atoi(sex)

		if err != nil {
			utils.Logger.Error("The sex fields must be number")
			c.JSON(http.StatusOK, gin.H{
				//Param error
				"err_code": 101,
				"msg":      "The sex param must be a number",
			})
			return
		}

		if sexInt == 1 {
			sex = "Male"
		} else if sexInt == 0 {
			sex = "Female"
		} else {
			sex = "Unknown"
		}

		//Send the message by slack
		msg := fmt.Sprintf("User IP:%s\nSex: %s\nContent:%s", userIP, sex, content)
		alert.SendMessage(msg)
		c.JSON(http.StatusOK, gin.H{
			//Param error
			"err_code": 0,
			"msg":      "Success",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		//Param error
		"err_code": 101,
		"msg":      "The content may be error",
	})
}
