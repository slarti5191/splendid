package utils

import (
	"github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"regexp"
	"time"
)

// SSHRunner holds the persistent pieces of a specific SSH device connection.
type SSHRunner struct {
	conn    *ssh.Client
	Ciphers []string
}

// Connect to the SSH Create the initial connection.
func (s *SSHRunner) Connect(user, pass, host string) {
	// Config for testing.
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// If we've been provided custom ciphers, add those to the config.
	if s.Ciphers != nil {
		config.Config = ssh.Config{
			Ciphers: s.Ciphers,
		}
	}

	// Dial up a connection.
	var err error
	s.conn, err = ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
}

// Gather sends commands over SSH and returns a match to m (regex) defined in our collector
func (s *SSHRunner) Gather(cmd []string, m *regexp.Regexp) string {
	// Spawn expect
	e, _, err := expect.SpawnSSH(s.conn, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to spawn expect: %v", err)
	}
	defer e.Close()

	// Send each command
	for c := range cmd {
		err := e.Send(cmd[c] + "\n")
		if err != nil {
			log.Fatalf("Error sending command %v: %v", c, err)
		}
	}

	// Wait for a match to "m" (regex passed from collector)
	_, match, err := e.Expect(m, 10*time.Second)
	if err != nil {
		log.Fatalf("Error matching config: %v", err)
	}

	// return config block
	return match[0]
}
