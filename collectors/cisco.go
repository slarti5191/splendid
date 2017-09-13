package collectors

import "fmt"

// devCisco lowercase means this is private to the collectors package.
type devCisco struct {

}

// Collect method is all that is necessary to implement the interface.
func (c devCisco) Collect() {
	fmt.Println("Running collect from CISCO!")
}

// MakeCisco creates a struct that implements the Collector
// interface for collecting Cisco configs.
func MakeCisco() Collector {
	return devCisco{}
}

// General command struct that may need to be refactored out.
type Commands struct {
	Commands map[string]string
}

// CiscoCmd generates a Commands struct with a list of strings
// of expected input and output.
func CiscoCmd() Commands {
	// Set commands to their expected output (last line)
	return Commands {
		map[string]string{
			"set pager": "",
			"show run":  "#",
		},
	}
}