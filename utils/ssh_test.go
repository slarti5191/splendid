package utils

import (
	"fmt"
	"github.com/slarti5191/splendid/configuration"
	"testing"
)

func TestSSHRunner_Connect(t *testing.T) {

}
func TestSSHRunner(t *testing.T) {
	// Pull in test.conf
	config, err := configuration.GetConfigs("../test.conf")
	if err != nil {
		t.Errorf("Problem: %s", err)
		// TODO: Once GetConfigs is fully tested, we might be able to skip this test.
		t.Skipf("test.conf not present")
	}
	fmt.Printf("User: %v\n", config.DefaultUser)

}
