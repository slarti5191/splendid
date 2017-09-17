package collectors

import "errors"

// Collector interface is the only common need between the various collectors.
// Ultimately, the core routine only cares to run collect on each collector.
// It doesn't need to know anything else about the inner implementation.
type Collector interface {
	Collect() string
}

// MakeCollector will generate the appropriate collector based on the
// type string passed in by the configuration.
func MakeCollector(m string) (Collector, error) {
	switch m {
	case "cisco_csb":
		return makeCiscoCsb(), nil
	case "cisco":
		return makeCisco(), nil
	case "pfsense":
		return makePfsense(), nil
	default:
		return nil, errors.New("unrecognized collector type")
	}
}
