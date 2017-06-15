package config

import (
	cg "github.com/olebedev/config"
	"io/ioutil"
	"os"
)

var Configure *cg.Config

func InitConfigInstance() error {
	path, _ := os.Getwd()
	cfgFile, err := ioutil.ReadFile(path + "/config/config.yaml")
	if err != nil {
		return err
	}
	cfg, err := cg.ParseYaml(string(cfgFile))
	if err != nil {
		return err
	}
	//根据当前环境变量获取配置
	env := os.Getenv("ENV")
	if len(env) == 0 {
		return err
	}

	Configure, err = cfg.Get(env)
	if err != nil {
		return err
	}

	return nil
}
