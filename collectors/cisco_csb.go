package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
	"regexp"
	"time"
)

type devCiscoCsb struct {
	configuration.DeviceConfig
}

func (d devCiscoCsb) Collect() string {
	// Set up SSH
	s := new(utils.SSHRunner)
	s.ReadWait = 2 * time.Second
	s.Ciphers = []string{"aes256-cbc", "aes128-cbc"}

	// Regex matching our config block
	var csb = regexp.MustCompile(`#([\s\S]*)#`)

	// Connect the SSHRunner
	s.Connect(d.User, d.Pass, d.Host)
	s.StartShell()
	defer s.Close()

	// Return our config
	return s.ShellCmd([]string{"terminal datadump", "show running-config", "exit"}, csb)
}

func makeCiscoCsb(d configuration.DeviceConfig) Collector {
	return &devCiscoCsb{d}
}
