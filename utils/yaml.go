package utils

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"server/global"
)

const configFile = "config.yaml"

func LoadYaml() ([]byte, error) {
	return os.ReadFile(configFile)
}

func SaveYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, byteData, fs.ModePerm)
}
