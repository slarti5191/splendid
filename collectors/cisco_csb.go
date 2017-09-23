package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"regexp"
)

type devCiscoCsb struct {
	configuration.DeviceConfig
}

func (d devCiscoCsb) Collect() string {
	// Regex matching our config block
	var csb = regexp.MustCompile(`#[\s\S]*?#`) // This likely doesn't work, untested regex
	// Commands we need to run
	cmd := []string{"terminal datadump", "show running-config", "exit"}
	// Set up SSH
	s := new(utils.SSHRunner)
	// Connect
	con := s.Connect(d.User, d.Pass, d.Host)
	// Return our config
	return s.Gather(con, cmd, csb)
}

func makeCiscoCsb(d configuration.DeviceConfig) Collector {
	return &devCiscoCsb{d}
}
