package configuration

import (
	"github.com/go-ini/ini"
	"log"
)

func mergeConfigFile(config *Config, filePath string) {
	iniFile, err := ini.Load(filePath)
	if err != nil {
		log.Println(filePath)
		panic(err) // Critical....
		//return nil, err
	}

	// Map [main] to config.
	err = iniFile.Section("main").MapTo(config)
	if err != nil {
		panic(err)
	}

	// TODO: Map/Load Device Configs
}
func getFileConfig(filePath string) Config {

	iniFile, err := ini.Load(filePath)
	if err != nil {
		log.Println(filePath)
		panic(err) // Critical....
		//return nil, err
	}

	// Map [main] to config.
	config := new(Config)
	err = iniFile.Section("main").MapTo(config)
	if err != nil {
		panic(err)
	}

	// TODO: Map/Load Device Configs
	return *config

	//return Config{
	//	false,
	//	"file.conf",
	//	88,
	//}
}
