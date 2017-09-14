package configuration

import (
	"errors"
	"fmt"
	"time"
)

type DeviceConfig struct {
	Hostname       string
	Method         string
	Target         string
	Timeout        time.Duration
	CommandTimeout time.Duration
	Config         map[string]string
}

type SplendidConfig struct {
	Interval      time.Duration
	Timeout       time.Duration
	GitPush       bool
	Insecure      bool
	Concurrency   int
	HttpListen    string
	HttpEnabled   bool
	SmtpString    string
	Workspace     string
	ExecutableDir string
	ToEmail       string
	FromEmail     string
	UseSyslog     bool
	DefaultUser   string
	DefaultPass   string
	DefaultMethod string
	CmwPass       string
	Devices       []DeviceConfig
	ConfigFile    string
}

// SetConfigs loads the config file, then parses flags
func SetConfigs() (*SplendidConfig, error) {
	Conf, err := LoadConfig()
	if err != nil {
		return nil, errors.New("Error loading configuration")
	}
	parseConfigFlags(Conf)
	fmt.Println(Conf)
	return Conf, nil
}
