package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"server/config"
	"server/utils"
)

func InitConf() *config.Config {
	c := &config.Config{}
	bytes, err := utils.LoadYaml()
	if err != nil {
		log.Fatalf("load yaml error: %v", err)
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		log.Fatalf("unmarshal yaml error: %v", err)
	}
	return c
}
