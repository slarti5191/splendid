package splendid

import "github.com/slarti5191/splendid/collectors"

// RunCollector collects configs
// Grab global configs as Conf, device specific commands as Cmd
func RunCollector(Dev DeviceConfig, Opts SplendidConfig, Cmds collectors.Commands) {
	// iterate over Cmds, expect matching output, fail otherwise

}
