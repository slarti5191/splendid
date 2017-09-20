package splendid

import (
	"fmt"
	"github.com/slarti5191/splendid/collectors"
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
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
	//s.config, err = configuration.GetConfigs("splendid.example.conf")
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
	// Build our collectors.
	s.cols = make([]collectors.Collector, 0, len(s.config.Devices))
	for _, c := range s.config.Devices {
		if c.Disabled {
			// Device config set to disabled=true to quickly turn off.
			log.Printf("Config disabled: %v", c.Name)
			continue
		}
		collector, err := collectors.MakeCollector(c)
		if err != nil {
			panic(err)
		}
		s.cols = append(s.cols, collector)
	}

	// Main collector loop.
	for {
		fmt.Println("> Running Collector Loop")

		// TODO: Ensure the below is running concurrently. Implement max concurrency setting.
		for _, c := range s.cols {
			go func(c collectors.Collector) {
				result := c.Collect()
				utils.WriteFile(result, c.GetName(), *s.config)
				log.Printf("Completed [%v] Len = %v", c.GetName(), len(result))
			}(c)
		}

		// Sleep until time for the next check.
		time.Sleep(s.config.Interval)
	}
}
