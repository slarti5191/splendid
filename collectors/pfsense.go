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
	var cmd []string
	// Regex matching our config block
	var pf = regexp.MustCompile(`<pfsense>[\s\S]*?<\/pfsense>`)
	// Commands we need to run
	// commands are different for "admin" user
	switch d.User {
	case "admin":
		cmd = append(cmd, "8", "cat /conf/config.xml", "exit", "0")
	default:
		cmd = append(cmd, "cat /conf/config.xml", "exit")
	}
	// Set up SSH
	s := new(utils.SSHRunner)
	// Connect
	s.Connect(d.User, d.Pass, d.Host)
	// Return our config
	return s.Gather(cmd, pf)
}

func makePfsense(d configuration.DeviceConfig) Collector {
	return &devPfsense{d}
}
