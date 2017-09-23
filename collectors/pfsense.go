package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"regexp"
)

type devPfsense struct {
	configuration.DeviceConfig
}

// Collect gathers config.xml from pfSense
func (d devPfsense) Collect() string {
	// Regex matching our config block
	var pf = regexp.MustCompile(`<pfsense>[\s\S]*?<\/pfsense>`)
	// Commands we need to run
	cmd := []string{"8", "cat /conf/config.xml", "exit", "0"}
	// Set up SSH
	s := new(utils.SSHRunner)
	// Connect
	con := s.Connect(d.User, d.Pass, d.Host)
	// Return our config
	return s.Gather(con, cmd, pf)
}

func makePfsense(d configuration.DeviceConfig) Collector {
	return &devPfsense{d}
}
