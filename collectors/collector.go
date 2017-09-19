package collectors

import (
	"errors"
	"github.com/slarti5191/splendid/configuration"
)

// Collector interface is the only common need between the various collectors.
// Ultimately, the core routine only cares to run collect on each collector.
// It doesn't need to know anything else about the inner implementation.
type Collector interface {
	Collect() string
	GetName() string
}

// MakeCollector will generate the appropriate collector based on the
// type string passed in by the configuration.
func MakeCollector(d configuration.DeviceConfig) (Collector, error) {
	switch d.Type {
	case "cisco_csb":
		return makeCiscoCsb(d), nil
	case "cisco":
		return makeCisco(d), nil
	case "pfsense":
		return makePfsense(d), nil
	default:
		return nil, errors.New("unrecognized collector type")
	}
}
