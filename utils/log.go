package utils

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/lndj/school-helper/config"
)

//The Logger instance
//Example: Logger.Warn("This is warn info")
var Logger *log.Logger

func init() {
	logFile, _ := config.Configure.String("log_file")
	if len(logFile) == 0 {
		appRoot, _ := os.Getwd()
		logFile = filepath.Join(appRoot, "app_runtime.log")
	}

	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Failed to log to file, using default stderr")
	}
	Logger = log.New()

	if env := config.Environment.AppEnv; env == "production" {
		//Log as JSON
		Logger.Formatter = &log.JSONFormatter{}
		//Output to file when production
		Logger.Out = f
		// Only log the debug severity or above.
		Logger.Level = log.WarnLevel
	} else {
		//Output to file when production
		Logger.Out = os.Stdout
		// Only log the debug severity or above.
		Logger.Level = log.DebugLevel
	}
}
