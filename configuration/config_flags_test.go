package configuration

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

// setOSFlagsForTesting provides a convenient reset for args and flags.
func setOSFlagsForTesting(args []string) {
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.CommandLine.Usage = func() { flag.Usage() }
}

// Basic test to ensure flags are processed.
func TestFlagsBasic(t *testing.T) {
	args := []string{
		"splendid",
		"-c", "../test.conf",
		"-p", "5",
		"-i", "75s",
		"-t", "85s",
		"-x=true",
		"--smtp", "smtp.example:3333",
		"-w=true",
		"--listen", "web.example:4444",
	}
	setOSFlagsForTesting(args)
	parseConfigFlags()

	// Ensure we can reset flags and test again.
	setOSFlagsForTesting(args)
	parseConfigFlags()

	// Check the flags that were found to be set.
	found := []string{}
	flag.Visit(func(flag *flag.Flag) {
		found = append(found, flag.Name)
	})
	// Flags sorts and calls "Visit" alphabetically.
	expect := []string{"c", "i", "listen", "p", "smtp", "t", "w", "x"}
	if !reflect.DeepEqual(found, expect) {
		t.Fatalf("Flags not parsed properly.\nFound: %s\nExpected: %s", found, expect)
	}
}

// TODO: Test for mergeConfigFlags and implement more flags.
