package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/nlopes/slack"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/lndj/school-helper/alert"
	"github.com/lndj/school-helper/config"
	"github.com/lndj/school-helper/router"
	"github.com/lndj/school-helper/utils"
)

const defaultPort = "8080"

func main() {
	defer closeLogFile()

	configRuntime()
	startSlackApp()
	alert.InitTelegram()
	startGin()
}

func configRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func startGin() {
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

//Start the SlackListener
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

//When exit, close the log file.
func closeLogFile() {
	if file, ok := utils.Logger.Out.(*os.File); ok {
		file.Sync()
		file.Close()
	} else if handler, ok := utils.Logger.Out.(io.Closer); ok {
		handler.Close()
	}
}
