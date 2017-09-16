package configuration

import (
	"time"
)

type DeviceConfig struct {
	Host           string
	Type           string
	Target         string
	User           string
	Pass           string
	Port           int
	Timeout        time.Duration
	CommandTimeout time.Duration
	Config         map[string]string
}

type SplendidConfig struct {
	Debug         bool
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

// GetConfigs loads the config file, then parses flags
func GetConfigs(configFile string) (*SplendidConfig, *DeviceConfig, error) {
	conf, dconf, err := loadConfig(configFile)
	if err != nil {
		return nil, nil, err
	}
	parseConfigFlags(conf)

	return conf, dconf, nil
}
