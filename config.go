package splendid

import (
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
}

// setConfigs parses splendid.conf
func SetConfigs() SplendidConfig {
	// set up config defaults
	Conf := *new(SplendidConfig)
	Conf.Workspace = "./splendid-workspace"
	Conf.Concurrency = 30
	Conf.SmtpString = "localhost:25"
	Conf.Interval = 300 * time.Second
	Conf.Timeout = 60 * time.Second
	Conf.Insecure = false
	Conf.GitPush = false
	Conf.HttpEnabled = false
	Conf.HttpListen = "localhost:5000"
	return Conf
}
