package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestMakeCisco(t *testing.T) {
	c := makeCisco(configuration.DeviceConfig{
		Name: "cisco",
		Host: "localhost",
		Type: "cisco",
		User: "user",
		Pass: "pass",
		Port: 22,
	})
	s := c.Collect()
	if s != "<xml>Example</xml>" {
		t.Errorf("Expected xml response missing.")
	}
}
