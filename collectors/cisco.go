package collectors

import "github.com/slarti5191/splendid/configuration"

// devCisco lowercase means this is private to the collectors package.
type devCisco struct {
	configuration.DeviceConfig
	commands [][]string
}

// Collect method is all that is necessary to implement the interface.
func (c devCisco) Collect() string {
	return "<xml>Example</xml>"
}

// makeCisco implements the Collector
// interface for collecting Cisco configs.
func makeCisco(d configuration.DeviceConfig) Collector {
	// Set commands to their expected output
	// each command gets a new slice containing the
	// command and a string expected after execution
	return devCisco{
		d,
		[][]string{
			{"set pager", ""},
			{"show run", "#"},
		},
	}
}
