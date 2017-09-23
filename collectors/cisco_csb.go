package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	//"github.com/slarti5191/splendid/utils"
)

type devCiscoCsb struct {
	configuration.DeviceConfig
}

func (d devCiscoCsb) Collect() string {
	//s := new(utils.SSHRunner)
	//s.Ciphers = []string{"aes256-cbc", "aes128-cbc"}

	//s.Connect(d.User, d.Pass, d.Host)
	//s.StartShell()

	// TODO: Parse version
	//version := s.ShellCmd("show version")
	//log.Printf("Ver: %v\n", version)

	//s.ShellCmd("terminal datadump")
	//result := s.ShellCmd("show running-config")
	//s.Close()

	return ""
}

func makeCiscoCsb(d configuration.DeviceConfig) Collector {
	return &devCiscoCsb{d}
}
