package configuration

import (
	"time"
)

// Config holds the general application settings.
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
	SMTPString   string
	ToEmail      string
	FromEmail    string

	// Webserver
	WebserverEnabled bool
	HTTPListen       string

	// Devices
	Devices []DeviceConfig
}

// DeviceConfig holds the device specific settings.
type DeviceConfig struct {
	Name           string
	Host           string
	Type           string
	User           string
	Pass           string
	Port           int
	Disabled       bool
	CustomTimeout  time.Duration
	CommandTimeout time.Duration
}

// GetName provides a simple implementation for the Collector interface.
func (d DeviceConfig) GetName() string {
	return d.Name
}

// getConfigDefaults provides reasonable defaults for Splendid!
func getConfigDefaults() *Config {
	return &Config{
		false,
		"splendid.conf",
		30,
		300 * time.Second,
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
