package configuration

import (
	"flag"
	"time"
)

// parseConfigFlags overwrites configuration with set flags
func parseConfigFlags(Conf *SplendidConfig) (*SplendidConfig, error) {
	// TODO: This should only happen if the flag has been passed
	flag.StringVar(&Conf.Workspace, "w", "./splendid-workspace", "Workspace")
	flag.IntVar(&Conf.Concurrency, "c", 30, "Number of collector processes")
	flag.StringVar(&Conf.SmtpString, "s", "localhost:25", "SMTP server:port")
	flag.DurationVar(&Conf.Interval, "interval", 300*time.Second, "Run interval")
	flag.DurationVar(&Conf.Timeout, "timeout", 60*time.Second, "Collection timeout")
	flag.BoolVar(&Conf.Insecure, "insecure", false, "Allow untrusted SSH keys")
	flag.BoolVar(&Conf.GitPush, "push", false, "Git push after commit")
	flag.BoolVar(&Conf.HttpEnabled, "web", false, "Run an HTTP status server")
	flag.StringVar(&Conf.HttpListen, "listen", "localhost:5000", "Host and port to use for HTTP status server (default: localhost:5000).")
	flag.StringVar(&Conf.ConfigFile, "f", "sample.conf", "Config File")
	flag.Parse()
	return Conf, nil
}
