package collectors

// devCisco lowercase means this is private to the collectors package.
type devCisco struct {
	commands map[string]string
}

// Collect method is all that is necessary to implement the interface.
func (c devCisco) Collect() string {
	return "<xml>Example</xml>"
}

// MakeCisco creates a struct that implements the Collector
// interface for collecting Cisco configs.
func makeCisco() Collector {
	// Set commands to their expected output (last line)
	return devCisco{
		map[string]string{
			"set pager": "",
			"show run":  "#",
		},
	}
}
