package splendid

import "github.com/slarti5191/splendid/collectors"

func Init() {
	// Get global configs
	Conf := SetConfigs()
	// Get device configs
	Dev := DeviceConfig{}
	// Need to come up with a way to set up 'Cmds'
	// as type Commands based on DeviceConfig.Method
	Cmds := splendid.CiscoCmd()

	// Kick off a collector
	RunCollector(&Dev, &Conf, Cmds)
}
