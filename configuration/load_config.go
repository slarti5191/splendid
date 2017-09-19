package configuration

import (
//"fmt"
//"github.com/go-ini/ini"
//"log"
)

// loadConfig loads the saved config file
// config is intentionally NOT passed by reference, so we can modify
// it without modifying the original defaults instance.
//func oldLoadConfig(configFile string, config SplendidConfig) (*SplendidConfig, error) {
//	// Load the INI file.
//	cfg, err := ini.Load(configFile)
//	if err != nil {
//		log.Println(configFile)
//		panic(err)
//		return nil, err
//	}
//
//	err = mainConfig(cfg.Section("main"), &config)
//	if err != nil {
//		return nil, err
//	}
//
//	// Grab device configs
//	err = devConfig(cfg, &config)
//	if err != nil {
//		return nil, err
//	}
//	return &config, err
//}
//
//// mainConfig maps the main section to SplendidConfig
//func mainConfig(cfg *ini.Section, conf *SplendidConfig) error {
//	err := cfg.MapTo(conf)
//	if err != nil {
//		return fmt.Errorf("error mapping main config: %v", err)
//	}
//	return nil
//}
//
//// devConfig maps the device config section to SplendidConfig.Devices
//func devConfig(cfg *ini.File, conf *SplendidConfig) error {
//	for _, b := range cfg.Sections() {
//		dconf := DeviceConfig{}
//		if b.HasKey("Host") {
//			b.MapTo(&dconf)
//			conf.Devices = append(conf.Devices, dconf)
//		}
//	}
//	return nil
//}
