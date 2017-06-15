package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
	"runtime"
	"school-helper/router"
	"school-helper/router/middleware"

	_ "school-helper/store"
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
	startGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func startGin() {
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
