package collectors

import "testing"

func TestMakeCisco(t *testing.T) {
	c := makeCisco()
	s := c.Collect()
	if s != "<xml>Example</xml>" {
		t.Errorf("Expected xml response missing.")
	}
}
