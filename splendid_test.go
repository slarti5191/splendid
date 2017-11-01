package splendid

import (
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestValidCollectorBuilds(t *testing.T) {
	s := Splendid{
		config: &configuration.Config{
			Devices: []configuration.DeviceConfig{{
				Name:           "pfsense",
				Host:           "localhost",
				Type:           "pfsense",
				User:           "user",
				Pass:           "pass",
				Port:           22,
				Disabled:       false,
				CustomTimeout:  30,
				CommandTimeout: 30,
			}},
		},
	}

	s.buildCollectors()

	// We expect this to work, but with zero resultant collectors.
	if len(s.cols) != 1 {
		t.Errorf("Expected 1 collectors. Got: %v", len(s.cols))
	}
}

func TestFakeCollectorDoesNotBuild(t *testing.T) {
	s := Splendid{
		config: &configuration.Config{
			Devices: []configuration.DeviceConfig{{
				Name:           "pfsense",
				Host:           "localhost",
				Type:           "fake",
				User:           "user",
				Pass:           "pass",
				Port:           22,
				Disabled:       false,
				CustomTimeout:  30,
				CommandTimeout: 30,
			}},
		},
	}

	s.buildCollectors()

	// We expect this to work, but with zero resultant collectors.
	if len(s.cols) != 0 {
		t.Errorf("Expected 0 collectors. Got: %v", len(s.cols))
	}
}
