package collectors

import (
	"github.com/slarti5191/splendid/configuration"
	"github.com/slarti5191/splendid/utils"
)

type devPfsense struct {
	configuration.DeviceConfig
}

func (d devPfsense) Collect() string {
	s := new(utils.SSHRunner)
	s.Connect(d.User, d.Pass, d.Host)

	// TODO: Non-primary user, press the "8" key.
	// Likely to need an SSH shell instead. Plus,
	// the shellRunner will need a different terminator. May
	// need the custom expect logic.
	result := s.Send("cat /conf/config.xml")
	s.Close()

	return result
}

func makePfsense(d configuration.DeviceConfig) Collector {
	return &devPfsense{d}
}
