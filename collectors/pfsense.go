package collectors

import (
	"github.com/slarti5191/splendid/utils"
)

type devPfsense struct {
}

func BuildSSHCollector(c devPfsense) *utils.SSHRunner {
	s := new(utils.SSHRunner)
	s.Connect()
	return s
}

func (c devPfsense) Collect() string {
	s := BuildSSHCollector(c)
	// TODO: Non-primary user, press the "8" key.
	result := s.Send("cat /conf/config.xml")
	return result
}

func makePfsense() Collector {
	return devPfsense{}
}
