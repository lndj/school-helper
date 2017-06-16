package main

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/gin-gonic/gin.v1"

	"school-helper/router"
)

const defaultPort = "8080"

func main() {
	ConfigRuntime()
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

//获取端口号
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	return port
}
