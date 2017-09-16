package configuration

import (
	"fmt"
	"github.com/go-ini/ini"
)

// loadConfig loads the saved config file
func loadConfig(configFile string) (*SplendidConfig, error) {
	conf := new(SplendidConfig)
	conf.ConfigFile = configFile
	// Load the INI file.
	cfg, err := ini.Load(conf.ConfigFile)
	if err != nil {
		return nil, err
	}

	c, err := mainConfig(cfg.Section("main"), conf)
	// Grab device configs
	err = devConfig(cfg, c)
	if err != nil {
		return nil, err
	}
	return c, err
}

// mainConfig maps the main section to SplendidConfig
func mainConfig(cfg *ini.Section, conf *SplendidConfig) (*SplendidConfig, error) {
	err := cfg.MapTo(conf)
	if err != nil {
		return nil, fmt.Errorf("error mapping main config: %v", err)
	}
	return conf, err
}

// devConfig maps the device config section to SplendidConfig.Devices
func devConfig(cfg *ini.File, conf *SplendidConfig) error {
	for _, b := range cfg.Sections() {
		dconf := DeviceConfig{}
		if b.HasKey("Host") {
			b.MapTo(&dconf)
			conf.Devices = append(conf.Devices, dconf)
		}
	}
	return nil
}
