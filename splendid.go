package splendid

import (
	"fmt"
	"github.com/slarti5191/splendid/collectors"
	"github.com/slarti5191/splendid/configuration"
	"time"
)

// Two primary threads. Webserver and collectors.
const version = "0.0.0"

func Init() {
	config, deviceconfig, err := configuration.GetConfigs("sample.conf")
	if err != nil {
		panic(err)
	}

	if config.Debug {
		fmt.Println("DEBUG ENABLED")
	}
	fmt.Println(config)
	fmt.Println(deviceconfig)

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
func runCollector(Dev configuration.DeviceConfig, Opts configuration.SplendidConfig) {
	// iterate over Cmds, expect matching output, fail otherwise
	fmt.Print(Dev, Opts)
}
