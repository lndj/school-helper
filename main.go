package main

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/nlopes/slack"

	"school-helper/router"
	"school-helper/alert"
)

const (
	defaultPort = "8080"

	botToken          = "xoxb-199470707778-xtG6Emt6iDttpCkK4XUJyVx0"
	botID             = "U5VDULTNW"
	clientSecret      = "2daa0356b329ac2bb889b40cb5c33dda"
	verificationToken = "Uc3kpfOLChVlR2n08Y2fikuj"
	channelID         = "C5VDCLP98"
)

func main() {
	ConfigRuntime()
	startSlackApp()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartGin() {
	env := os.Getenv("ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Load all the router
	r := router.Load()

	fmt.Printf("Run as [%s] environment\n", env)
	fmt.Printf("Listen and serve on %s\n", port())
	r.Run(":" + port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	return port
}

func startSlackApp() {
	client := slack.New(botToken)
	slackListener := &alert.SlackListener{
		Client:    client,
		BotID:     botID,
		ChannelID: channelID,
	}
	alert.SL = slackListener
	fmt.Printf("SlackListener is running on channel:%s\n", channelID)
	go slackListener.ListenAndResponse()
}
