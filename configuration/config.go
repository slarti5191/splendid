package configuration

import (
	"log"
	"reflect"
	"time"
	"fmt"
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
func GetConfigs(configFile string) (*SplendidConfig, error) {
	// 0: Initialize Defaults
	defaults := SplendidConfig{
		false,
		30,
		120,
		false,
		false,
		30,
		"localhost:5001",
		true,
		"server:port",
		"/workspace",
		"/",
		"",
		"",
		false,
		"user",
		"pass",
		"none",
		"none",
		nil,
		"sample.conf",
	}

	if configFile != "" {
		defaults.ConfigFile = configFile
	}

	// 1) Need to load flags first to determine which config file to use.
	flags := parseConfigFlags(defaults)
log.Println(flags)
	// 2) Load in config from file on top of defaults array.
	config, err := loadConfig(flags.ConfigFile, defaults)
	if err != nil {
		return nil, fmt.Errorf("Error[%s] %s", flags.ConfigFile, err)
	}

	// 3) If flag value is provided by user, apply override.
	config.flagUpdate(flags, defaults)

	// 4) And convert to seconds where needed.
	config.Interval = config.Interval * time.Second
	config.Timeout = config.Timeout * time.Second

	return config, nil
}

func (c *SplendidConfig) flagUpdate(flags SplendidConfig, defaults SplendidConfig) {

	//e, d, err := Compare(flags, defaults)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(e)
	//log.Println(d)
	//os.Exit(0)

	defaultValue := reflect.ValueOf(defaults)
	//loadedValue := reflect.ValueOf(*c)
	flagValue := reflect.ValueOf(flags)

	configValue := reflect.ValueOf(c).Elem()

	for i := 0; i < flagValue.NumField(); i++ {
		// Grab default value and flag value.
		v1 := defaultValue.Field(i)
		v2 := flagValue.Field(i)

		// Unwraps pointers if necessary.
		v1 = reflect.Indirect(v1)
		v2 = reflect.Indirect(v2)

		// We don't support structs and slices via flag input.
		if v1.Type().Kind() == reflect.Struct || v1.Type().Kind() == reflect.Slice {
			continue
		}

		// Simple sanity checks.
		if v1.Type() != v2.Type() {
			log.Fatal("Expected types to match.")
		}
		if !v1.IsValid() || !v2.IsValid() {
			log.Fatal("Expect valid values.")
		}

		// Check if something was set.
		// TODO: Problem, explicitly set command line values that match default
		// values are not picked up. Deferring to the config file value.
		if v1.Interface() != v2.Interface() {
			// If we have a zero value for the flag, skip it.
			if v2.Interface() == reflect.Zero(v2.Type()).Interface() {
				continue
			}

			//log.Printf("Override: %v", defaultValue.Type().Field(i).Name)
			//log.Printf("%v -->> %v", v1.Interface(), v2.Interface())

			// Override the config value with flag value.
			configValue.Field(i).Set(v2)
		}
	}
}
