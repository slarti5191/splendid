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
		"-x=true", // Explicitly pass true.
		"-dc",     // Implicitly pass true.
		"-e",
		"-w",
		"--listen", "web.example:4444",
		"--copyrights",
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
	expect := []string{"c", "copyrights", "dc", "e", "i", "listen", "p", "t", "w", "x"}
	if !reflect.DeepEqual(found, expect) {
		t.Fatalf("Flags not parsed properly.\nFound: %s\nExpected: %s", found, expect)
	}
}

// Test merging a value for every config flag.
func TestMergeAllConfigFlags(t *testing.T) {
	setOSFlagsForTesting([]string{
		"splendid",
		"-c", "../test.conf",
		"-p", "5",
		"-i", "75s",
		"-t", "85s",
		"--debug",
		"-x",
		"-dc",
		"-e",
		"-w",
		"--listen", "web.example:4444",
		"--copyrights",
	})

	parseConfigFlags()
	config := new(Config)
	mergeConfigFlags(config)
}
