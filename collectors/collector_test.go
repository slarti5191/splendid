package collectors

import (
	"testing"
)

func TestMakeCollector(t *testing.T) {
	c, err := MakeCollector("cisco")
	if err != nil {
		t.Errorf("Making Cisco should not error.")
	}
	_, ok := c.(devCisco)
	if !ok {
		t.Errorf("Expected type not received.")
	}

	c, err = MakeCollector("fake")
	if err == nil {
		t.Errorf("Expected an error for a fake collector type.")
	}
	if c != nil {
		t.Errorf("Expected a nil result for fake collector type.")
	}
}
