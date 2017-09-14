package configuration

import (
	"testing"
)

func testConfiguration(t *testing.T) {
	configFile := "sample.conf"
	_, err := GetConfigs(configFile)
	if err != nil {
		t.Errorf("Configs should not error.")
	}
}
