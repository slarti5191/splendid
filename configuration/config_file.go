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

	// Map/Load Device Configs
	for _, section := range iniFile.Sections() {
		if section.Name() == "main" || section.Name() == "DEFAULT" {
			continue
		}
		if !section.HasKey("Type") {
			log.Panicf("Device [%v] must have a type defined.", section.Name())
		}
		deviceConfig := DeviceConfig{}
		section.MapTo(&deviceConfig)
		config.Devices = append(config.Devices, deviceConfig)
	}
}

//
//func getFileConfig(filePath string) Config {
//
//	iniFile, err := ini.Load(filePath)
//	if err != nil {
//		log.Println(filePath)
//		panic(err) // Critical....
//		//return nil, err
//	}
//
//	// Map [main] to config.
//	config := new(Config)
//	err = iniFile.Section("main").MapTo(config)
//	if err != nil {
//		panic(err)
//	}
//
//	// TODO: Map/Load Device Configs
//	return *config
//
//	//return Config{
//	//	false,
//	//	"file.conf",
//	//	88,
//	//}
//}
