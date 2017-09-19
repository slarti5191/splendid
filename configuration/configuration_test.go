package configuration

import (
	"flag"
	"os"
	"testing"
)

//func TestConfiguration(t *testing.T) {
//	// Need to overwrite flags....
//	os.Args = []string{"splendid", "-x"}
//
//	configFile := "../test.conf"
//	//ResetForTesting(nil)
//	config, err := GetConfigs(configFile)
//	if err != nil {
//		t.Error(err)
//		t.Fatalf("Open: %v Err: %v", configFile, err)
//		t.Fatalf("Open: %v Err: %v", configFile, err)
//	}
//	//https://golang.org/src/flag/flag_test.go
//	//ResetForTesting(nil)
//	//_, _ = GetConfigs(configFile)
//
//	// Expects
//	if config.Timeout != 120*time.Second {
//		t.Fatal("Expected: 30 Got: %v", config.Timeout)
//	}
//	if config.DefaultUser != "splendid" {
//		t.Fatal("Expected: splendid Got: %v", config.DefaultUser)
//	}
//}

func xTestOverrides(t *testing.T) {

}

func ResetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.Usage = func() { flag.Usage() }
	flag.Usage = usage
}

func xTestFlags(t *testing.T) {
	os.Args = []string{"splendid", "-f", "../splendid.conf"}
	//ResetForTesting(nil)
	//sc := SplendidConfig{}
	//_ = parseConfigFlags(sc)

	//flagSet := flag.CommandLine
	//args := []string{
	//	"-f", "../test.conf",
	//}
	//conf := flagSet.String("g", "splendid.example.conf", "Sample conf file.")
	//if err := flagSet.Parse(args); err != nil {
	//	t.Fatal(err)
	//}
	if !flag.Parsed() {
		t.Error("flagSet.Parsed() = false after Parse")
	}
	//configs, _ := GetConfigs("")
	//if configs == nil {
	//	t.Error("Nil config?")
	//}
	//if *conf != "splendid.conf" {
	//	t.Errorf("Error: %s", *conf)
	//}
}
