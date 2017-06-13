package config

import (
	"github.com/labstack/gommon/log"
	configer "github.com/olebedev/config"
)

func Get(env string) interface{} {
	cfg, err := configer.ParseYaml("./config.yaml")

	if err != nil {
		log.Fatalf("Parse yaml file error:%v", err)

		return nil
	}

	content, err := cfg.Get(env)

	if err != nil {
		log.Fatalf("Get config error:%v", err)
		return nil
	}

	return content
}
