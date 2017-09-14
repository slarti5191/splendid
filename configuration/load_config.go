package configuration

import (
	"errors"
	"github.com/go-ini/ini"
)

// LoadConfig loads the saved config file
func LoadConfig() (*SplendidConfig, error) {
	Conf := new(SplendidConfig)
	// Load config file
	Conf.ConfigFile = "sample.conf"
	cfg, err := ini.Load(Conf.ConfigFile)
	if err != nil {
		return nil, errors.New("Unable to open config file")
	}

	Conf.DefaultUser = cfg.Section("main").Key("DefaultUser").Value()
	return Conf, nil
}
