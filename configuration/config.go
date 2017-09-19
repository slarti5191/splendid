package configuration

import (
	"log"
	"reflect"
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

// GetConfigs loads the config file, then parses flags
//func GetConfigs(configFile string) (*SplendidConfig, error) {
//	// 0: Initialize Defaults
//	defaults := SplendidConfig{
//		false,
//		30,
//		120,
//		false,
//		false,
//		30,
//		"localhost:5001",
//		true,
//		"server:port",
//		"/workspace",
//		"/",
//		"",
//		"",
//		false,
//		"user",
//		"pass",
//		"none",
//		"none",
//		nil,
//		"splendid.example.conf",
//	}
//
//	if configFile != "" {
//		defaults.ConfigFile = configFile
//	}
//
//	// 1) Need to load flags first to determine which config file to use.
//	flags := parseConfigFlags(defaults)
//	log.Println(flags)
//
//	// ....
//	// TODO: Bad...
//	if flags.ConfigFile == "" {
//		flags.ConfigFile = defaults.ConfigFile
//	}

// 2) Load in config from file on top of defaults array.
//config, err := loadConfig(flags.ConfigFile, defaults)
//if err != nil {
//	panic(err)
//	return nil, fmt.Errorf("Error[%s] %s", flags.ConfigFile, err)
//}
//
//// 3) If flag value is provided by user, apply override.
//config.flagUpdate(flags, defaults)
//
//// 4) And convert to seconds where needed.
//config.Interval = config.Interval * time.Second
//config.Timeout = config.Timeout * time.Second

//return config, nil
//}

// mergeConfig was an attempt to merge one config into another config.
// However, it's not very useful since all it really does is merge everything
// into the first config. There is not a way to say "if this value is not set..."
// because sometimes you WANT to merge a false or a zero value into the config.
// So there is no way to distinguish a "nil" unset value (ZeroValue) and a
// desired zero value...
func (c *Config) mergeConfig(merge Config) {

	mergeValue := reflect.ValueOf(merge)
	//curValue := reflect.ValueOf(*c)
	target := reflect.ValueOf(c).Elem()

	for i := 0; i < mergeValue.NumField(); i++ {
		// Grab default value and flag value.
		cur := target.Field(i)
		val := mergeValue.Field(i)

		// Unwraps pointers if necessary.
		//cur = reflect.Indirect(cur)
		//val = reflect.Indirect(val)

		// We don't support structs and slices via flag input.
		if val.Type().Kind() == reflect.Struct || val.Type().Kind() == reflect.Slice {
			continue
		}

		// Simple sanity checks.
		if !val.IsValid() {
			log.Fatal("Expect valid values.")
		}

		if cur.Interface() != val.Interface() {
			// If we have a zero value for the flag, skip it.
			//if val.Interface() == reflect.Zero(val.Type()).Interface() {
			//continue
			//}

			//log.Printf("Override: %v", defaultValue.Type().Field(i).Name)
			//log.Printf("%v -->> %v", val.Interface(), v2.Interface())

			// Override the config value with flag value.
			target.Field(i).Set(val)
		}
	}
}

// Keeping the below for reference for the moment.
//func (c *SplendidConfig) flagUpdate(flags SplendidConfig, defaults SplendidConfig) {
//
//	//e, d, err := Compare(flags, defaults)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//log.Println(e)
//	//log.Println(d)
//	//os.Exit(0)
//
//	defaultValue := reflect.ValueOf(defaults)
//	//loadedValue := reflect.ValueOf(*c)
//	flagValue := reflect.ValueOf(flags)
//
//	configValue := reflect.ValueOf(c).Elem()
//
//	for i := 0; i < flagValue.NumField(); i++ {
//		// Grab default value and flag value.
//		v1 := defaultValue.Field(i)
//		v2 := flagValue.Field(i)
//
//		// Unwraps pointers if necessary.
//		v1 = reflect.Indirect(v1)
//		v2 = reflect.Indirect(v2)
//
//		// We don't support structs and slices via flag input.
//		if v1.Type().Kind() == reflect.Struct || v1.Type().Kind() == reflect.Slice {
//			continue
//		}
//
//		// Simple sanity checks.
//		if v1.Type() != v2.Type() {
//			log.Fatal("Expected types to match.")
//		}
//		if !v1.IsValid() || !v2.IsValid() {
//			log.Fatal("Expect valid values.")
//		}
//
//		// Check if something was set.
//		// TODO: Problem, explicitly set command line values that match default
//		// values are not picked up. Deferring to the config file value.
//		if v1.Interface() != v2.Interface() {
//			// If we have a zero value for the flag, skip it.
//			if v2.Interface() == reflect.Zero(v2.Type()).Interface() {
//				continue
//			}
//
//			//log.Printf("Override: %v", defaultValue.Type().Field(i).Name)
//			//log.Printf("%v -->> %v", v1.Interface(), v2.Interface())
//
//			// Override the config value with flag value.
//			configValue.Field(i).Set(v2)
//		}
//	}
//}
