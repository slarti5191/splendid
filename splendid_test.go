package splendid

import (
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := SetConfigs()
	if err != nil {
		t.Errorf("Configs should not error.")
	}
}
