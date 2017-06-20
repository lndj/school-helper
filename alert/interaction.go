package alert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nlopes/slack"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/lndj/school-helper/config"
)

//InteractionHandler, the slack interaction api handler
func InteractionHandler(c *gin.Context) {
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("[ERROR] Failed to read request body: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		fmt.Printf("[ERROR] Failed to unespace request body: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		fmt.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	// Only accept message from slack with valid token
	if message.Token != config.Environment.SlackVerifyToken {
		fmt.Printf("[ERROR] Invalid token: %s", message.Token)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
		return
	}

	action := message.Actions[0]
	switch action.Name {
	case actionSelect:
		value := action.SelectedOptions[0].Value
		// Overwrite original drop down message.
		originalMessage := message.OriginalMessage
		originalMessage.Attachments[0].Text = fmt.Sprintf("Confirm to run [%s] ?", strings.Title(value))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionStart,
				Text:  "Yes",
				Type:  "button",
				Value: "start",
				Style: "primary",
			},
			{
				Name:  actionCancel,
				Text:  "No",
				Type:  "button",
				Style: "danger",
			},
		}

		c.JSON(http.StatusOK, originalMessage)
	case actionStart:

		//Run the task, Just a example
		go func() {
			fmt.Printf("[INFO] Start to run the task\n")
			time.Sleep(2 * time.Second)
			fmt.Printf("[INFO] Task is end\n")

			//Send the result
			SendMessage("Your task run success!")
		}()

		title := ":ok: Your task is running! yay!"
		responseMessage(c, message.OriginalMessage, title, "")
	case actionCancel:
		title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
		responseMessage(c, message.OriginalMessage, title, "")
	default:
		fmt.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
	}

}

//Response to slack
func responseMessage(c *gin.Context, original slack.Message, title, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}

	c.JSON(http.StatusOK, original)
}
