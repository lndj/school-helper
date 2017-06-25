package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/lndj/school-helper/config"
	"github.com/lndj/school-helper/utils"
)

//InteractionHandler, the slack interaction api handler
func InteractionHandler(c *gin.Context) {
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Logger.Errorf("Failed to read request body: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		utils.Logger.Errorf("Failed to unespace request body:%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		utils.Logger.Errorf("Failed to decode json message from slack:%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	// Only accept message from slack with valid token
	if message.Token != config.Environment.SlackVerifyToken {
		utils.Logger.Errorf("Invalid token:%s", message.Token)
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
				Value: value,
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
		if len(action.Value) > 0 {
			go runCommand(action.Value)
		}
		title := ":ok: Your task is running! yay!"
		responseMessage(c, message.OriginalMessage, title, "")
	case actionCancel:
		title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
		responseMessage(c, message.OriginalMessage, title, "")
	default:
		utils.Logger.Errorf("Invalid action was submitted:%s", action.Name)
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

func runCommand(command string) error {

	utils.Logger.Infof("Start to run the task, the command is:%s", command)

	cmd := exec.Command(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		utils.Logger.Fatal(err)
	}

	var content string
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(3 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			utils.Logger.Fatal("Failed to kill: ", err)
		}
		content = out.String()
		utils.Logger.Println("Process killed as timeout reached")
	case err := <-done:
		if err != nil {
			utils.Logger.Printf("Process done with error = %v", err)
		} else {
			content = out.String()
			utils.Logger.Print("Process done gracefully without error")
		}
	}
	content = "Your task run success!\n```" + content + "```"
	SendMessage(content)
	utils.Logger.Info("Task is end")
	return nil
}
