package configuration

import (
	"flag"
)

// parseConfigFlags reads in configuration with set flags
func parseConfigFlags(flags *SplendidConfig, defaults SplendidConfig) {
	// Set to passed flags, otherwise go with config
	flag.IntVar(&flags.Concurrency, "c", defaults.Concurrency, "Number of collector processes")
	flag.StringVar(&flags.SmtpString, "s", defaults.SmtpString, "SMTP")
	//flag.DurationVar(&flags.Interval, "interval", defaults.Interval, "Run interval")
	//flag.DurationVar(&flags.Timeout, "timeout", defaults.Timeout, "Collection timeout")
	flag.BoolVar(&flags.Insecure, "insecure", defaults.Insecure, "Allow untrusted SSH keys")
	flag.BoolVar(&flags.HttpEnabled, "web", defaults.HttpEnabled, "Run an HTTP status server")
	flag.StringVar(&flags.HttpListen, "listen", defaults.HttpListen, "Host and port to use for HTTP status server.")
	flag.StringVar(&flags.ConfigFile, "f", defaults.ConfigFile, "Config File")
	flag.Parse()
}
