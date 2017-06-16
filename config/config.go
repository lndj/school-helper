package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	cg "github.com/olebedev/config"
)

var Configure *cg.Config

func init() {
	appRoot, _ := os.Getwd()
	configFile := filepath.Join(appRoot, "config/config.yaml")
	cfgFile, err := ioutil.ReadFile(configFile)
	checkErr(err)

	cfg, err := cg.ParseYaml(string(cfgFile))
	checkErr(err)

	//Get config by environment
	env := os.Getenv("ENV")
	if len(env) == 0 {
		panic(err)
	}

	Configure, err = cfg.Get(env)
	//Parse the Env into config
	Configure.Env()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
