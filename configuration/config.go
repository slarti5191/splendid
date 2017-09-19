package configuration

import (
	"time"
)

type Config struct {
	// Debug
	Debug bool

	// Config
	ConfigFile string

	// Collector Settings
	Concurrency      int
	Interval         time.Duration
	Timeout          time.Duration
	AllowInsecureSSH bool

	// Git
	GitPush bool

	// User Settings
	Workspace   string
	DefaultUser string
	DefaultPass string

	// Mail
	EmailEnabled bool
	SmtpString   string
	ToEmail      string
	FromEmail    string

	// Webserver
	WebserverEnabled bool
	HttpListen       string

	// Devices
	Devices []DeviceConfig
}

// Do we need these in Config?
// UseSyslog bool
// ExecutableDir string
// DefaultMethod string
// CmwPass       string

type DeviceConfig struct {
	Host           string
	Type           string
	User           string
	Pass           string
	Port           int
	CustomTimeout  time.Duration
	CommandTimeout time.Duration
}

// What is DeviceConfig.Target used for?
// Target         string

func getConfigDefaults() *Config {
	return &Config{
		false,
		"splendid.conf",
		30,
		30 * time.Second,
		120 * time.Second,
		false,
		false,
		"/workspace",
		"",
		"",
		false,
		"smtp:port",
		"",
		"",
		false,
		"localhost:5002",
		nil,
	}
}
