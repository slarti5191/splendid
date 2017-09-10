package main

import (
	"github.com/slarti5191/splendid"
)
const version = "0.0.0"

// main initializes the configuration
// and kicks off the collectors
func main() {
	Conf := splendid.SetConfigs()
	Dev := splendid.DeviceConfig{}
	// Obviously this doesn't work, but we
	// should initialize commands based on DeviceConfig method
	Cmd := Dev.Method
	// set Cmd to match the appropriate method for the device
	// this may not be the right place for this?
	splendid.RunCollector(&Conf, Cmd)

}
