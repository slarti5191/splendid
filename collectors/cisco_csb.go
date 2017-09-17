package collectors

import (
	"github.com/slarti5191/splendid/utils"
	"log"
)

type devCiscoCsb struct {
}

func (d devCiscoCsb) Collect() string {
	//ssh.Config -> Ciphers: []string{"aes128-cbc", "aes256-cbc", "none"},
	s := new(utils.SSHRunner)
	s.Ciphers = []string{"aes256-cbc", "aes128-cbc"}

	log.Println("Attempting to connect...")
	s.Connect("splendid", "Splendid1", "switch.lan.hdthings.com")
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

func makeCiscoCsb() Collector {
	return new(devCiscoCsb)
}
