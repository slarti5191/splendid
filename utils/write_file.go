package utils

import (
	"github.com/slarti5191/splendid/configuration"
	"os"
)

func WriteFile(c string, name string, s configuration.Config) {
	var configPath string
	configPath = s.Workspace + "/" + name
	config, err := os.Create(configPath)
	if err != nil {
		panic(err)
	}
	defer config.Close()
	config.WriteString(c)
	config.Sync()
}
