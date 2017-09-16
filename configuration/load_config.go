package configuration

import (
	"fmt"
	"github.com/go-ini/ini"
)

// loadConfig loads the saved config file
func loadConfig(configFile string) (*SplendidConfig, *DeviceConfig, error) {
	conf := new(SplendidConfig)
	conf.ConfigFile = configFile
	// Load the INI file.
	cfg, err := ini.Load(conf.ConfigFile)
	if err != nil {
		return nil, nil, err
	}

	mainconfig, err := mainConfig(cfg.Section("main"), conf)
	if err != nil {
		return nil, nil, err
	}

	deviceconfig, err := devConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	return mainconfig, deviceconfig, err
}

// mainConfig maps the main section to SplendidConfig
func mainConfig(cfg *ini.Section, conf *SplendidConfig) (*SplendidConfig, error) {
	err := cfg.MapTo(conf)
	if err != nil {
		return nil, fmt.Errorf("error mapping main config: %v", err)
	}

	return conf, err
}

// devConfig maps the device config section to DeviceConfig
func devConfig(cfg *ini.File) (*DeviceConfig, error) {
	dconf := new(DeviceConfig)
	// TODO: This only works with one device - last device in the config is set here.
	for _, b := range cfg.Sections() {
		if b.HasKey("Host") {
			err := b.MapTo(dconf)
			if err != nil {
				return nil, err
			}
		}
	}
	
	return dconf, nil
}
