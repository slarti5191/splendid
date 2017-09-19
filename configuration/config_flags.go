package configuration

import (
	"flag"
	"log"
)

func parseConfigFlags() {
	defaults := getConfigDefaults()

	flag.String("c", defaults.ConfigFile, "Config file name.")
	flag.Int("p", defaults.Concurrency, "Parallel concurrent threads to use for collection.")

	// Interval and Timeout
	flag.Duration("i", defaults.Interval, "Interval in seconds between run calls.")
	flag.Duration("t", defaults.Timeout, "Timeout default in seconds to wait for collection to finish.")

	flag.Bool("debug", defaults.Debug, "Enable DEBUG flag for development.")
	flag.Bool("x", defaults.AllowInsecureSSH, "Allow untrusted SSH keys.")

	flag.String("smtp", defaults.SmtpString, "SMTP connection string.")

	flag.Bool("w", defaults.WebserverEnabled, "Run a web status server.")
	flag.String("listen", defaults.HttpListen, "Host and port to use for HTTP status server.")

	flag.Parse()
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

func mergeConfigFlags(config *Config) {
	flag.Visit(func(flagVal *flag.Flag) {
		switch flagVal.Name {
		case "c":
			config.ConfigFile = flagVal.Value.(flag.Getter).Get().(string)
			//config.ConfigFile = string(flag.Value)
		case "p":
			config.Concurrency = flagVal.Value.(flag.Getter).Get().(int)
		case "debug":
			config.Debug = flagVal.Value.(flag.Getter).Get().(bool)
			// TODO: MORE!
		default:
			log.Fatalf("Flag merge not configured for %v", flagVal)
		}
	})
}
