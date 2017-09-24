package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"regexp"
	"time"
)

type devPfsense struct {
	configuration.DeviceConfig
}

// Collect gathers config.xml from pfSense
func (d devPfsense) Collect() string {
	// Set up SSH
	s := new(utils.SSHRunner)
	// PFSense seems to return everything in one blast.
	// 10ms, 100ms is too quick, does not pull anything in some tests.
	// TODO: Make it more intelligent to wait till it receives the first
	// bit due to latency, then secondary blurbs use ReadWait.
	// With 250ms, a network hiccup could return an empty string.
	s.ReadWait = 250 * time.Millisecond

	// Regex matching our config block
	var pf = regexp.MustCompile(`<pfsense>[\s\S]*?<\/pfsense>`)

	// Connect the SSHRunner
	s.Connect(d.User, d.Pass, d.Host)
	s.StartShell()
	defer s.Close()

	return s.ShellCmd([]string{"8", "cat /conf/config.xml", "exit", "0"}, pf)
}

func makePfsense(d configuration.DeviceConfig) Collector {
	return &devPfsense{d}
}
