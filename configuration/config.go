package configuration

import (
	"time"
	"reflect"
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
	c, dconf, err := loadConfig(configFile)
	if err != nil {
		return nil, nil, err
	}
	flags, err := parseConfigFlags(c)
	conf := c.flagUpdate(*flags)
	return &conf, dconf, nil
}

func (c SplendidConfig) flagUpdate(f SplendidConfig) (conf SplendidConfig) {
	old := reflect.ValueOf(c)
	new := reflect.ValueOf(f)
	final := reflect.ValueOf(&conf).Elem()
	for i := 0; i < old.NumField(); i++ {
		if !new.Field(i).IsValid() {
			final.Field(i).Set(new.Field(i))
		} else {
			final.Field(i).Set(old.Field(i))
		}
	}
	return
}