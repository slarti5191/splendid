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
	s := new(utils.SSHRunner)
	s.Ciphers = []string{"aes256-cbc", "aes128-cbc"}
	// Regex matching our config block
	var csb = regexp.MustCompile(`#[\s\S]*?#`) // This likely doesn't work, untested regex
	// Commands we need to run
	cmd := []string{"terminal datadump", "show running-config", "exit"}
	// Set up SSH
	// Connect
	s.Connect(d.User, d.Pass, d.Host)
	// Return our config
	s.StartShell()
	return s.ShellCmd(cmd, csb)
	// s.Gather depends on google/expect which is not cross platform
	//return s.Gather(cmd, csb)
}

func makeCiscoCsb(d configuration.DeviceConfig) Collector {
	return &devCiscoCsb{d}
}
