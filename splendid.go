package splendid

import (
	"fmt"
	"github.com/slarti5191/splendid/collectors"
	"time"
)

// Two primary threads. Webserver and collectors.
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
	Conf := SetConfigs()
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

	// Main collector loop.
	for {
		fmt.Println("> Running Collector Loop")

		// Kick off a collector
		RunCollector(Dev, Conf, Cmds)

		time.Sleep(5 * time.Second)
	}
}
