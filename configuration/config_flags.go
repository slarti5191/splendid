package configuration

import (
	"flag"
	"time"
)

// parseConfigFlags overwrites configuration with set flags
func parseConfigFlags(Conf *SplendidConfig) (*SplendidConfig, error) {
	// Set to passed flags, otherwise go with config
	flag.IntVar(&Conf.Concurrency, "c", Conf.Concurrency, "Number of collector processes")
	flag.StringVar(&Conf.SmtpString, "s", Conf.SmtpString, "SMTP server:port")
	flag.DurationVar(&Conf.Interval, "interval", Conf.Interval*time.Second, "Run interval")
	flag.DurationVar(&Conf.Timeout, "timeout", Conf.Timeout*time.Second, "Collection timeout")
	flag.BoolVar(&Conf.Insecure, "insecure", Conf.Insecure, "Allow untrusted SSH keys")
	flag.BoolVar(&Conf.HttpEnabled, "web", Conf.HttpEnabled, "Run an HTTP status server")
	flag.StringVar(&Conf.HttpListen, "listen", Conf.HttpListen, "Host and port to use for HTTP status server (default: localhost:5000).")
	flag.StringVar(&Conf.ConfigFile, "f", Conf.ConfigFile, "Config File")
	flag.Parse()
	return Conf, nil
}
