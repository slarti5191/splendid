package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

type SSHRunner struct {
	conn *ssh.Client
	sess *ssh.Session
}

func (s *SSHRunner) Connect() {
	// Config for testing.
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("pass"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// Dial up a connection.
	var err error
	s.conn, err = ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	// Prematurely closes the connection when the function loses scope.
	//defer conn.Close()

	// Start a new session for the connection.
	s.sess, err = s.conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
}

// Send the command and receive the result.
func (s *SSHRunner) Send(cmd string) string {
	// Execute the command via ssh, and return the output.
	out, err := s.sess.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	return string(out)
}

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
