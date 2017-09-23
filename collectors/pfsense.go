package collectors

import (
	"github.com/slarti5191/splendid/configuration"
//	"github.com/slarti5191/splendid/utils"
	"github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
	"net"
	"regexp"
)

type devPfsense struct {
	configuration.DeviceConfig
}

func (d devPfsense) Collect() string {
	//s := new(utils.SSHRunner)
	//con := s.Connect(d.User, d.Pass, d.Host)

	config := ssh.ClientConfig{
		User: d.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(d.Pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// If we've been provided custom ciphers, add those to the config.
/*	if s.Ciphers != nil {
		config.Config = ssh.Config{
			Ciphers: s.Ciphers,
		}
	}*/

	var pf = regexp.MustCompile(`<pfsense>[\s\S]*?<\/pfsense>`)
	// Dial up a connection.
	var err error
	conn, err := ssh.Dial("tcp", d.Host+":22", &config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	e, _, err := expect.SpawnSSH(conn, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to spawn Expect")
	}
	defer e.Close()
	if err := e.Send("8" + "\n"); err != nil {
		log.Fatal(err)
	}
	if err := e.Send("cat /conf/config.xml" + "\n"); err != nil {
		log.Fatal(err)
	}
	_, match, err  := e.Expect(pf, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	e.Send("exit" + "\n")
	e.Send("0" + "\n")
	//fmt.Println(out)
	ret := match[0]
	//if b != nil {
	//	log.Fatalf("Expect not running")
	//}
	//fmt.Println(a)
	// TODO: Non-primary user, press the "8" key.
	// Likely to need an SSH shell instead. Plus,
	// the shellRunner will need a different terminator. May
	// need the custom expect logic.
	//result := s.Send("cat /conf/config.xml")
	//s.Close()

	return ret
}

func makePfsense(d configuration.DeviceConfig) Collector {
	return &devPfsense{d}
}
