package utils

import (
	"testing"
)

// TODO: This is going to require a little finesse.
// We are unable to use functions from other test packages.
// We can pull in setOSFlagsForTesting, but then we also
// need to figure out how to reset config loader singleton.
// That's why a singleton is often an anti-pattern, due to
// the headaches of testing.
// Once we figure out that, we should set up a localtest.conf
// file (gitignore) that this test can take advantage of to
// actually run ssh tests against live pfsense/cisco devices.
func TestSSHRunner(t *testing.T) {
	// Pull in test.conf
	//_, err := configuration.GetConfigs("../test.conf")
	//if err != nil {
	//	t.Skipf("test.conf not present")
	//}
}
