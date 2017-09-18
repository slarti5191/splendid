package configuration

import (
	"flag"
	"log"
	"time"
)

func fetchDefaults() SplendidConfig {
	return SplendidConfig{
		false,
		30 * time.Second,
		120 * time.Second,
		false,
		false,
		30,
		"localhost:5001",
		false,
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
}

func init() {
	log.Println("=-- FLAGS: INIT --=")
	// Runs before unit tests. :(
	createFlags()
}

func createFlags() {
	log.Println("=-- FLAGS: Creating --=")
	defaults := fetchDefaults()

	flag.String("c", defaults.ConfigFile, "Config file name.")
	flag.Int("p", defaults.Concurrency, "Parallel concurrent threads to use for collection.")

	// Interval and Timeout
	flag.Duration("i", defaults.Interval, "Interval in seconds between run calls.")
	flag.Duration("t", defaults.Timeout, "Timeout default in seconds to wait for collection to finish.")

	flag.Bool("x", defaults.Insecure, "Allow untrusted SSH keys.")

	flag.String("smtp", defaults.SmtpString, "SMTP connection string.")

	flag.Bool("w", defaults.HttpEnabled, "Run a web status server.")
	flag.String("listen", defaults.HttpListen, "Host and port to use for HTTP status server.")
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

	//setta := func(n string, v flag.Value) {
	//	switch n {
	//	case "c":
	//		flags.Concurrency, _ = strconv.Atoi(v.String())
	//	case "f":
	//		flags.ConfigFile = v.String()
	//	}
	//}
	//
	//flag.Visit(func(flag *flag.Flag) {
	//	log.Printf("Flag[%s] %s", flag.Name, flag.Value)
	//	setta(flag.Name, flag.Value)
	//})

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