package configuration

import (
	"testing"
	"time"
)

func TestConfiguration(t *testing.T) {
	configFile := "../test.conf"
	config, err := GetConfigs(configFile)
	if err != nil {
		t.Fatalf("getconfigs returned err: %v", err)
	}

	// Expects
	if config.Timeout != 120*time.Second {
		t.Fatal("Expected: 30 Got: %v", config.Timeout)
	}
	if config.DefaultUser != "splendid" {
		t.Fatal("Expected: splendid Got: %v", config.DefaultUser)
	}
}
