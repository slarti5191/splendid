package configuration

import (
	"testing"
)

func testConfiguration(t *testing.T) {
	_, err := GetConfigs()
	if err != nil {
		t.Errorf("Configs should not error.")
	}
}
