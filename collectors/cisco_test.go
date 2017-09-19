package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestMakeCisco(t *testing.T) {
	c := makeCisco(configuration.DeviceConfig{
		"localhost",
		"cisco",
		"user",
		"pass",
		22,
		30,
		30,
	})
	s := c.Collect()
	if s != "<xml>Example</xml>" {
		t.Errorf("Expected xml response missing.")
	}
}
