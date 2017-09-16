package utils

import (
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestSSHRunner_Connect(t *testing.T) {

}
func TestSSHRunner(t *testing.T) {
	// Pull in test.conf
	_, _, err := configuration.GetConfigs("../test.conf")
	if err != nil {
		t.Skipf("test.conf not present")
	}
}
