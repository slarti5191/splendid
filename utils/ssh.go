package utils

import (
	"golang.org/x/crypto/ssh"
	"log"
	"fmt"
	"net"
)

func Test() {
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("pass"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	conn, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	sess, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	// Example file for cisco.
	out, err := sess.CombinedOutput("cat /conf/config.xml")
	if err != nil {
		panic(err)
	}
	// Currently just dumps to command line, but can save string(out) to file.
	fmt.Println(string(out))
}
