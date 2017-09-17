package utils

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"sync"
)

type SSHRunner struct {
	conn     *ssh.Client
	sess     *ssh.Session
	shellIn  chan<- string
	shellOut <-chan string
	Ciphers  []string
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

	// Prematurely closes the connection when the function loses scope.
	//defer conn.Close()

	// Start a new session for the connection.
	s.sess, err = s.conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
}

// StartShell must be called once before running ShellCmd.
func (s *SSHRunner) StartShell() {
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := s.sess.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatal(err)
	}

	w, err := s.sess.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := s.sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	in, out := shellRunner(w, r)
	//if err := session.Start("/bin/sh"); err != nil {
	if err := s.sess.Shell(); err != nil {
		log.Fatal(err)
	}
	// Ignore initially connection text.
	<-out

	s.shellIn = in
	s.shellOut = out
}

// ShellCmd sends a request to the SSH connection and returns the output.
func (s *SSHRunner) ShellCmd(cmd string) string {
	if s.shellIn == nil {
		log.Fatal("Shell not yet initialized.")
	}
	s.shellIn <- cmd
	result := <-s.shellOut
	//fmt.Printf("version: %s\n", result)
	//fmt.Println("=-----------------------------=")
	return result
}

// shellRunner is an internal function used by SSHRunner to manage the input
// and output of a shell.
func shellRunner(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			t += n
			//log.Println(t)
			//log.Println(string(buf[:t]))
			//if t > 0 {
			//	log.Println(string(buf[t-1]))
			//}
			// This is essentially "expect" pulling in the buffer
			// until the expected string is found. In this case
			// this is currently only used for cisco_csb which has
			// a hash symbol for the prompt. Like 'switch#'
			// Alternative would use t-2 and $ if the prompt is
			// something like the $PS1 == 'sh-4.3$ '
			if t > 0 && buf[t-1] == '#' {
				//log.Printf("Returning: %v", string(buf[:t]))
				// TODO: Trim off the shell prompt...
				// Might be easier to trim with the expect logic.
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}

// Close attempts to close the session and connection.
func (s *SSHRunner) Close() {
	s.sess.Close()
	s.conn.Close()
}

// Send a single command for the session, returns the result.
// If more than one command needs to be run, a shell must be used.
func (s *SSHRunner) Send(cmd string) string {
	// Execute the command via ssh, and return the output.
	//log.Printf("Run: %v", cmd)

	out, err := s.sess.CombinedOutput(cmd)
	if err != nil {
		if err.Error() == "ssh: Stdout already set" {
			log.Fatalf("Send function can not be run twice")
		}
		//log.Println("Errored...")
		//log.Println(err.Error())
		panic(err)
	}
	//log.Println("Completed!")
	//log.Println(string(out))

	return string(out)
}
