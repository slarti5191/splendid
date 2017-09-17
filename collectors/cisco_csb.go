package collectors

import (
	"github.com/slarti5191/splendid/utils"
	"log"
	"github.com/slarti5191/splendid/configuration"
)

type devCiscoCsb struct {
	config configuration.DeviceConfig
}

func (d devCiscoCsb) Collect() string {
	s := new(utils.SSHRunner)
	s.Ciphers = []string{"aes256-cbc", "aes128-cbc"}

	log.Println("Attempting to connect...")
	s.Connect(d.config.User, d.config.Pass, d.config.Host)
	s.StartShell()

	log.Println("Connected, showing version...")
	version := s.ShellCmd("show version")
	log.Printf("Ver: %v\n", version)

	s.ShellCmd("terminal datadump")
	result := s.ShellCmd("show running-config")
	//log.Printf("Config: %v\n", result)

	s.Close()
	log.Printf("Closed")

	return result
}

func makeCiscoCsb(d configuration.DeviceConfig) Collector {
	return &devCiscoCsb{d}
}
