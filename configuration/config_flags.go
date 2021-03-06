package configuration

import (
	"flag"
	"log"
	"time"
)

func parseConfigFlags() {
	defaults := getConfigDefaults()

	flag.String("c", defaults.ConfigFile, "Config file name.")
	flag.Int("p", defaults.Concurrency, "Parallel concurrent threads to use for collection.")
	flag.Duration("i", defaults.Interval, "Interval in seconds between run calls.")
	flag.Duration("t", defaults.Timeout, "Timeout default in seconds to wait for collection to finish.")
	flag.Bool("debug", defaults.Debug, "Enable DEBUG flag for development.")
	flag.Bool("x", defaults.AllowInsecureSSH, "Allow untrusted SSH keys.")
	flag.Bool("e", defaults.EmailEnabled, "Enable or disable email when changes found.")
	flag.Bool("dc", defaults.DisableCollection, "Disable collector processing.")
	flag.Bool("w", defaults.WebserverEnabled, "Run a web status server.")
	flag.String("listen", defaults.HTTPListen, "Host and port to use for HTTP status server.")
	flag.Bool("copyrights", defaults.Copyrights, "Display copyright licenses of compiled packages.")

	flag.Parse()
}

// mergeConfigFlags maps the flag values back onto Config.
// There ought to be a more efficient way to handle this when
// combined with the above function for defining and parsing
// the flags. However, it is not apparent how to easily tell whether
// a flag was explicitly set to a default value or not. Plus,
// some other edge case considerations.
func mergeConfigFlags(config *Config) {
	flag.Visit(func(flagVal *flag.Flag) {
		switch flagVal.Name {
		// These should be in the same order that the flags above are declared.
		case "c":
			config.ConfigFile = flagVal.Value.(flag.Getter).Get().(string)
		case "p":
			config.Concurrency = flagVal.Value.(flag.Getter).Get().(int)
		case "i":
			config.Interval = flagVal.Value.(flag.Getter).Get().(time.Duration)
		case "t":
			config.Timeout = flagVal.Value.(flag.Getter).Get().(time.Duration)
		case "debug":
			config.Debug = flagVal.Value.(flag.Getter).Get().(bool)
		case "dc":
			config.DisableCollection = flagVal.Value.(flag.Getter).Get().(bool)
		case "x":
			config.AllowInsecureSSH = flagVal.Value.(flag.Getter).Get().(bool)
		case "e":
			config.EmailEnabled = flagVal.Value.(flag.Getter).Get().(bool)
		case "w":
			config.WebserverEnabled = flagVal.Value.(flag.Getter).Get().(bool)
		case "listen":
			config.HTTPListen = flagVal.Value.(flag.Getter).Get().(string)
		case "copyrights":
			config.Copyrights = flagVal.Value.(flag.Getter).Get().(bool)
		// Fail if not defined.
		default:
			log.Fatalf("Flag merge not configured for %v", flagVal)
		}
	})
}

func configFlagsGetConfigPath() string {
	// Flags must be parsed.
	if !flag.Parsed() {
		panic("Flags not yet parsed.")
	}

	// Grab the default config path.
	configPath := getConfigDefaults().ConfigFile

	// And check to see if we want to override it with a flag.
	configFlag := flag.Lookup("c")
	if configFlag != nil {
		found := flag.Lookup("c").Value.(flag.Getter).Get().(string)
		if found != configPath {
			//log.Printf("Config: %s, Switching to found: %s", configPath, found)
			configPath = found
		}
	}

	return configPath
}
