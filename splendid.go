package splendid

import (
	"github.com/slarti5191/splendid/collectors"
)

const version = "0.0.0"

func Init() {
	// Get global configs
	Conf := *SetConfigs()
	// Set up DeviceConfig
	Dev := new(DeviceConfig)
	Dev.Method = "cisco"
	// Get Commands out of collectors.Generate
	Cmds := collectors.Generate(string(Dev.Method))
	// Kick off a collector
	RunCollector(Dev, Conf, Cmds)
}
