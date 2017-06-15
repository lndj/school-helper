package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/labstack/gommon/log"
	"gopkg.in/gin-gonic/gin.v1"

	"school-helper/router"
	"school-helper/router/middleware"

	"school-helper/config"
	"school-helper/store"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
)

func main() {
	ConfigRuntime()
	InitConfig()
	InitStore()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

//Init the Configure instance
func InitConfig() {
	err := config.InitConfigInstance()
	if err != nil {
		log.Fatalf("Init config error: %v", err)
	}
}

//Init Storage instance
func InitStore() {
	//redis
	store.InitRedisClient()
}

func StartGin() {
	env := os.Getenv("ENV")
	if env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.Load(middleware.InitRedis())

	r.Run(":" + port()) // listen and serve on 0.0.0.0:8080
}

//获取端口号
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	return port
}
