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
	//Conf := *SetConfigs()
	// Get device configs
	//Dev := DeviceConfig{}

	// Set up DeviceConfig
	//Dev := new(DeviceConfig)
	//Dev.Method = "cisco"
	// Get Commands out of collectors.Generate
	//Cmds := collectors.Generate(string(Dev.Method))
	// Kick off a collector
	//RunCollector(Dev, Conf, Cmds)

	// TODO: Move this into the master collector.
	exampleCollector, err := collectors.MakeCollector("pfsense")
	if err != nil {
		fmt.Println("Oh no, an error")
	}

	// Main collector loop.
	for {
		fmt.Println("> Running Collector Loop")

		result := exampleCollector.Collect()
		fmt.Println(result)

		// Kick off a collector
		//runCollector(Dev, Conf)

		time.Sleep(5 * time.Second)
	}
}

// RunCollector collects configs
// Grab global configs as Conf, device specific commands as Cmd
func runCollector(Dev DeviceConfig, Opts SplendidConfig) {
	// iterate over Cmds, expect matching output, fail otherwise
	fmt.Print(Dev, Opts)
}
