package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nlopes/slack"
	"gopkg.in/gin-gonic/gin.v1"

	"school-helper/alert"
	"school-helper/router"
	"school-helper/config"
)

const defaultPort = "8080"

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
	port := config.Environment.AppPort
	if len(port) == 0 {
		port = defaultPort
	}
	return port
}

func startSlackApp() {
	client := slack.New(config.Environment.SlackBotToken)
	slackListener := &alert.SlackListener{
		Client:    client,
		BotID:     config.Environment.SlackBotId,
		ChannelID: config.Environment.SlackChannelID,
	}
	//Create a slack client instance
	alert.SL = slackListener
	fmt.Printf("SlackListener is running on channel:%s\n", config.Environment.SlackChannelID)
	go slackListener.ListenAndResponse()
}
