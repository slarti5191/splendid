package splendid

import (
	"fmt"
	"github.com/slarti5191/splendid/collectors"
	"github.com/slarti5191/splendid/configuration"
	"log"
	"os"
	"time"
)

// Two primary threads. Webserver and collectors.
const version = "0.0.0"

type Splendid struct {
	config *configuration.Config
	cols   []collectors.Collector
}

// Run is the entry point to the application.
func (s *Splendid) Run() {
	s.config = configuration.GetConfig()

	//var err error
	//s.config, err = configuration.GetConfigs("sample.conf")
	//if err != nil {
	//	panic(err)
	//}
	//
	if s.config.Debug {
		log.Println("DEBUG ENABLED: Dumping config and exiting.")
		log.Println(s.config)
		os.Exit(0)
	}

	// Kickstart the webserver if enabled.
	if s.config.WebserverEnabled {
		go s.threadWebserver()
	}
	// Kickstart the main collector thread.
	s.threadCollectors()
}

// threadWebserver is a placeholder for what will someday be a webserver.
func (s *Splendid) threadWebserver() {
	loopDelay := 5 * time.Second
	for {
		fmt.Println("> Webserver code on another branch.")
		time.Sleep(loopDelay)
		loopDelay *= 5
	}
}

// threadCollectors iterates through all device configs and runs the collectors.
func (s *Splendid) threadCollectors() {
	s.cols = make([]collectors.Collector, len(s.config.Devices))
	for i, c := range s.config.Devices {
		collector, err := collectors.MakeCollector(c)
		if err != nil {
			panic(err)
		}
		s.cols[i] = collector
	}

	// Main collector loop.
	for {
		fmt.Println("> Running Collector Loop")

		// TODO: Use concurrency!
		for _, c := range s.cols {
			go func() {
				result := c.Collect()
				fmt.Println(result)
			}()
		}

		time.Sleep(10 * time.Second)
	}
}

// RunCollector collects configs
// Grab global configs as Conf, device specific commands as Cmd
func runCollector(Dev configuration.DeviceConfig, Opts configuration.SplendidConfig) {
	// iterate over Cmds, expect matching output, fail otherwise
	fmt.Print(Dev, Opts)
}
