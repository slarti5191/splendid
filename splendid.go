package splendid

import (
	"github.com/slarti5191/splendid/collectors"
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"github.com/slarti5191/splendid/web"
	"log"
	"os"
	"sync"
	"time"
)

// Two primary threads. Webserver and collectors.
const version = "0.0.0"

// Splendid is the main container used to run the application.
type Splendid struct {
	config *configuration.Config
	git    *utils.Git
	cols   []collectors.Collector
}

// Run is the entry point to the application.
func (s *Splendid) Run() {
	s.config = configuration.GetConfig()

	if s.config.Debug {
		log.Println("DEBUG ENABLED: Dumping config and exiting.")
		log.Println(s.config)
		os.Exit(0)
	}

	// Initialize GIT
	s.git = &utils.Git{
		Path: s.config.Workspace,
	}
	err := s.git.Open()
	if err != nil {
		log.Fatalf("Could not open GIT repo: %v", err)
	}

	// Kickstart the webserver if enabled.
	if s.config.WebserverEnabled {
		if s.config.DisableCollection {
			// Collection is disabled, run on the main thread.
			web.RunTheServer()
		} else {
			// Start coroutine for webserver.
			go web.RunTheServer()
		}
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

	// If no collectors were built, we have a problem.
	if len(s.cols) == 0 {
		log.Fatal("Must set DisableCollection or pass -dc flag if no collectors are enabled.")
	}

	// Main collector loop.
	var waitGroup sync.WaitGroup
	for {
		log.Printf("> Running %v Collector(s)\n", len(s.cols))

		for _, c := range s.cols {
			waitGroup.Add(1)
			go s.runCollector(c, &waitGroup)
		}
		waitGroup.Wait()
		log.Println("> Devices collected. Processing diffs.")

		// TODO: Email the changes.
		_, err := s.git.GitCommit()
		if err != nil {
			panic(err)
		}

		// Sleep until time for the next check.
		log.Printf("> Collection routine complete. Next run in %v\n", s.config.Interval)
		time.Sleep(s.config.Interval)
	}
}

var failCounts = make(map[collectors.Collector]int)

func (s *Splendid) runCollector(c collectors.Collector, waitGroup *sync.WaitGroup) {
	log.Printf("Starting [%v]", c.GetName())

	result := c.Collect()
	if result == "" {
		// Happens sometimes... do not write an empty file.
		log.Printf("No result: [%v] was empty.\n", c.GetName())

		// Track fails for this collector.
		failCounts[c]++
		if failCounts[c] > 2 {
			// It appears we have a problem. (Email?)
			log.Fatalf("[%v] Failed three times in a row.", c.GetName())
		}
	} else {
		// Ensure reset of fail counts for this collector.
		failCounts[c] = 0
		utils.WriteFile(result, c.GetName(), *s.config)
		log.Printf("Completed [%v] Len = %v\n", c.GetName(), len(result))
	}

	waitGroup.Done()
}

func (s *Splendid) runCollectorGit(c collectors.Collector) {
	log.Printf("Starting [%v]", c.GetName())
	result := c.Collect()

	utils.WriteFile(result, c.GetName(), *s.config)
	log.Printf("Completed [%v] Len = %v", c.GetName(), len(result))
	//diff := s.git.GitHash(c.GetName())
	diff := s.git.GitDiff(c.GetName())
	//diff := s.git.GitDiff("test")
	if diff != "" {
		log.Printf("Discovered a change:\n%v\n", diff)
	}
}
