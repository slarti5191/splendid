package configuration

import (
	"flag"
	"log"
	"strconv"
)

func init() {
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

	flag.Int("c", defaults.Concurrency, "Number of collector processes")
	flag.String( "s", defaults.SmtpString, "SMTP")
	flag.Duration( "interval", defaults.Interval, "Run interval")
	flag.Duration("timeout", defaults.Timeout, "Collection timeout")
	flag.Bool("insecure", defaults.Insecure, "Allow untrusted SSH keys")
	flag.Bool( "web", defaults.HttpEnabled, "Run an HTTP status server")
	flag.String("listen", defaults.HttpListen, "Host and port to use for HTTP status server.")
	flag.String("f", defaults.ConfigFile, "Config File")
}

// https://stackoverflow.com/questions/17412908/how-do-i-unit-test-command-line-flags-in-go

//func defineFlags() {
//	//f := flag.NewFlagSet("General", flag.ContinueOnError)
//	flag.Parse([])
//}
var myFlags SplendidConfig

// parseConfigFlags reads in configuration with set flags
func parseConfigFlags(defaults SplendidConfig) SplendidConfig {
	flags := SplendidConfig{}

	flag.Parse()

	setta := func(n string, v flag.Value) {
		switch n {
		case "c":
			flags.Concurrency, _ = strconv.Atoi(v.String())
		case "f":
			flags.ConfigFile = v.String()
		}
	}

	flag.Visit(func(flag *flag.Flag) {
		log.Printf("Flag[%s] %s", flag.Name, flag.Value)
		setta(flag.Name, flag.Value)
	})

	return flags

	//if flag.Parsed() {
	//	log.Println("==--==-- ALREADY PARSED --==--==")
	//} else {
	//	flags := SplendidConfig{}
	//	// Set to passed flags, otherwise go with config
	//	flag.IntVar(&flags.Concurrency, "c", defaults.Concurrency, "Number of collector processes")
	//	flag.StringVar(&flags.SmtpString, "s", defaults.SmtpString, "SMTP")
	//	flag.DurationVar(&flags.Interval, "interval", defaults.Interval, "Run interval")
	//	flag.DurationVar(&flags.Timeout, "timeout", defaults.Timeout, "Collection timeout")
	//	flag.BoolVar(&flags.Insecure, "insecure", defaults.Insecure, "Allow untrusted SSH keys")
	//	flag.BoolVar(&flags.HttpEnabled, "web", defaults.HttpEnabled, "Run an HTTP status server")
	//	flag.StringVar(&flags.HttpListen, "listen", defaults.HttpListen, "Host and port to use for HTTP status server.")
	//	flag.StringVar(&flags.ConfigFile, "f", defaults.ConfigFile, "Config File")
	//	flag.Parse()
	//
	//	myFlags = flags
	//	log.Println("==--==-- FIRST PARSE --==--==")
	//	log.Println(myFlags)
	//	return flags
	//}
}
