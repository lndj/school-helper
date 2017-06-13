package config

import (
	"github.com/labstack/gommon/log"
	"github.com/murlokswarm/errors"
	cg "github.com/olebedev/config"
	"io/ioutil"
	"os"
)

type Config struct {
	Content *cg.Config
}

func NewConfig() (*cg.Config, error) {
	path, _ := os.Getwd()
	cfgFile, err := ioutil.ReadFile(path + "/config/config.yaml")
	if err != nil {
		log.Fatalf("Read yaml file error:%v", err)
		return nil, errors.New("Read yaml file error:%v", err)
	}

	cfg, err := cg.ParseYaml(string(cfgFile))

	if err != nil {
		log.Fatalf("Parse yaml file error:%v", err)
		return nil, errors.New("Parse yaml file error:%v", err)
	}
	cfg.Env()

	//根据当前环境变量获取配置
	env, err := cfg.String("ENV")

	if err != nil {
		log.Fatalf("Get env error:%v", err)
		return nil, errors.New("Get ENV error:%v", err)
	}

	ct := new(Config)
	ct.Content, err = cfg.Get(env)

	if err != nil {
		log.Fatalf("Get config error:%v", err)
		return nil, errors.New("Get config error:%v", err)
	}

	return ct.Content, nil
}
