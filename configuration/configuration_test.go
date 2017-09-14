package configuration

import (
	"testing"
)

func testConfiguration(t *testing.T) {
	_, err := SetConfigs()
	if err != nil {
		t.Errorf("Configs should not error.")
	}
}
