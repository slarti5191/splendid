package configuration

import (
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

// SetConfigs loads the config file, then parses flags
func GetConfigs(configFile string) (*SplendidConfig, error) {
	conf, err := loadConfig(configFile)
	if err != nil {
		// Just pass the err up the stack?
		return nil, err
		//return nil, errors.New("Error loading configuration")
	}
	parseConfigFlags(conf)

	if conf.Debug {
		fmt.Println("DEBUG ENABLED")
	}
	fmt.Println(conf)

	return conf, nil
}
