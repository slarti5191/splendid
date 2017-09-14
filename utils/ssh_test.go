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
		// TODO: Once GetConfigs is fully tested, we might be able to skip this test.
		//t.Errorf("Problem: %s", err)
		// However, we need to skip since travis fails otherwise.
		// Maybe create a test.conf file for travis?
		t.Skipf("test.conf not present")
	}
	fmt.Printf("User: %v\n", config.DefaultUser)

}
