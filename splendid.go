package splendid

import (
	"fmt"
	"github.com/slarti5191/splendid/collectors"
	"time"
)

// Two primary threads. Webserver and collectors.
const version = "0.0.0"

func Init() {
	go threadWebserver()
	threadCollectors()
}

func threadWebserver() {
	for {
		fmt.Println("> Webserver code on another branch.")
		time.Sleep(3 * time.Second)
	}
}

func threadCollectors() {
	// Get global configs
	Conf := *SetConfigs()
	// Get device configs
	Dev := DeviceConfig{}

	// Set up DeviceConfig
	//Dev := new(DeviceConfig)
	//Dev.Method = "cisco"
	// Get Commands out of collectors.Generate
	//Cmds := collectors.Generate(string(Dev.Method))
	// Kick off a collector
	//RunCollector(Dev, Conf, Cmds)

	// Need to come up with a way to set up 'Cmds'
	// as type Commands based on DeviceConfig.Method
	Cmds := collectors.CiscoCmd()

	// TODO: Move this into the master collector.
	exampleCollector, err := collectors.MakeCollector("cisco")
	if err != nil {
		fmt.Println("Oh no, an error")
	}

	// Main collector loop.
	for {
		fmt.Println("> Running Collector Loop")

		exampleCollector.Collect()

		// Kick off a collector
		runCollector(Dev, Conf, Cmds)

		time.Sleep(5 * time.Second)
	}
}

// RunCollector collects configs
// Grab global configs as Conf, device specific commands as Cmd
func runCollector(Dev DeviceConfig, Opts SplendidConfig, Cmds collectors.Commands) {
	// iterate over Cmds, expect matching output, fail otherwise
	fmt.Print(Dev, Opts, Cmds)
}
