package splendid

import (
	"flag"
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
func SetConfigs() *SplendidConfig {
	// set up config defaults
	Conf := new(SplendidConfig)
	//parse cli flags
	flag.StringVar(&Conf.Workspace, "w", "./splendid-workspace", "Workspace")
	flag.IntVar(&Conf.Concurrency, "c", 30, "Number of collector processes")
	flag.StringVar(&Conf.SmtpString, "s", "localhost:25", "SMTP server:port")
	flag.DurationVar(&Conf.Interval, "interval", 300*time.Second, "Run interval")
	flag.DurationVar(&Conf.Timeout, "timeout", 60*time.Second, "Collection timeout")
	flag.BoolVar(&Conf.Insecure, "insecure", false, "Allow untrusted SSH keys")
	flag.BoolVar(&Conf.GitPush, "push", false, "Git push after commit")
	flag.BoolVar(&Conf.HttpEnabled, "web", false, "Run an HTTP status server")
	flag.StringVar(&Conf.HttpListen, "listen", "localhost:5000", "Host and port to use for HTTP status server (default: localhost:5000).")
	flag.Parse()
	return Conf
}
