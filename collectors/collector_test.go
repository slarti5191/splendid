package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestMakeCollector(t *testing.T) {
	c, err := MakeCollector(configuration.DeviceConfig{
		"localhost",
		"cisco",
		"",
		"user",
		"pass",
		22,
		30,
		30,
	})
	if err != nil {
		t.Errorf("Making Cisco should not error.")
	}
	_, ok := c.(devCisco)
	if !ok {
		t.Errorf("Expected type not received.")
	}

	c, err = MakeCollector(configuration.DeviceConfig{
		"localhost",
		"fake",
		"",
		"user",
		"pass",
		22,
		30,
		30,
	})
	if err == nil {
		t.Errorf("Expected an error for a fake collector type.")
	}
	if c != nil {
		t.Errorf("Expected a nil result for fake collector type.")
	}
}

func TestMakeCollector2(t *testing.T) {
	c, err := MakeCollector(configuration.DeviceConfig{
		"localhost",
		"cisco",
		"",
		"user",
		"pass",
		22,
		30,
		30,
	})
	if err != nil {
		t.Errorf("Making Cisco should not error.")
	}
	_, ok := c.(devCisco)
	if !ok {
		t.Errorf("Expected type not received.")
	}

	c, err = MakeCollector(configuration.DeviceConfig{
		"localhost",
		"fake",
		"",
		"user",
		"pass",
		22,
		30,
		30,
	})
	if err == nil {
		t.Errorf("Expected an error for a fake collector type.")
	}
	if c != nil {
		t.Errorf("Expected a nil result for fake collector type.")
	}
}
