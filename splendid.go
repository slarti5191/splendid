package splendid

import (
	"github.com/slarti5191/splendid/collectors"
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"github.com/slarti5191/splendid/web"
	"log"
	"os"
	"time"
)

// Two primary threads. Webserver and collectors.
const version = "0.0.0"

// Splendid is the main container used to run the application.
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
		go web.RunTheServer()
	}
	// Kickstart the main collector thread.
	s.threadCollectors()
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
		// Silence if all collectors are disabled.
		if len(s.cols) > 0 {
			log.Printf("> Running %v Collector(s)", len(s.cols))
		}

		for _, c := range s.cols {
			go func(c collectors.Collector) {
				//log.Printf("Starting [%v]", c.GetName())
				result := c.Collect()
				utils.WriteFile(result, c.GetName(), *s.config)
				log.Printf("Completed [%v] Len = %v", c.GetName(), len(result))
			}(c)
		}

		// Sleep until time for the next check.
		time.Sleep(s.config.Interval)
	}
}
