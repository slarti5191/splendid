package utils

import (
	//"github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"regexp"
	"time"
)

// SSHRunner holds the persistent pieces of a specific SSH device connection.
type SSHRunner struct {
	conn     *ssh.Client
	sess     *ssh.Session
	shellIn  chan<- string
	shellOut <-chan string
	Ciphers  []string
	ReadWait time.Duration
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
	// Start a new session for the connection.
	s.sess, err = s.conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
}

// StartShell must be called once before running ShellCmd.
func (s *SSHRunner) StartShell() {
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,      // disable echoing
		ssh.TTY_OP_ISPEED: 115200, // input speed = console speed
		ssh.TTY_OP_OSPEED: 115200, // output speed = console speed
		//ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		//ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
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
	in, out := s.shellRunner(w, r)
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
func (s *SSHRunner) ShellCmd(cmd []string, re *regexp.Regexp) string {
	var result string
	if s.shellIn == nil {
		log.Fatal("Shell not yet initialized.")
	}
	for c := range cmd {
		//log.Printf("Run: %s\n======\n", cmd[c])
		s.shellIn <- cmd[c] + "\n"
		next := <-s.shellOut
		//log.Printf("Aggregate: %s\n======\n", next)
		result += next
	}

	// Run shellOut through a regex to pull the config
	return match(result, re)
}

// match parses input and returns a matching substring
func match(conf string, re *regexp.Regexp) string {
	//r := regexp.MustCompile(`<pfsense>[\s\S]*?<\/pfsense>`)
	result := re.FindString(conf)
	return (result)
}

// shellRunner is an internal function used by SSHRunner to manage the input
// and output of a shell.
func (s *SSHRunner) shellRunner(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	go func() {
		for cmd := range in {
			// Appending a + "\n" here does not yield the expected newline character.
			w.Write([]byte(cmd))
		}
	}()

	go func() {
		var (
			t      int
			result string
		)
		for {
			str := s.readDelayed(r)
			n := len(str)
			result += str
			t += n
			//log.Printf("Read: [+%v / %v]\n", n, t)
			if n == 0 {
				//log.Printf("ReadReturn: %v\n", result)
				out <- result
				// Reset running result.
				result = ""
				t = 0
			}
		}
	}()
	return in, out
}

// ReadDelayed will return an empty string if no result is returned within 3 seconds.
// TODO: Make the number of seconds based on the device config.
// TODO: Integrate with SSHRunner and keep a persistent buf[] since we may be losing
// some of the input when we timeout.
// TODO: Attempted to move "buf" to the SSHRunner, but realized "n" also needs to be
// moved since the function returns while the previous read is blocked.
func (s *SSHRunner) readDelayed(r io.Reader) string {
	out := make(chan string, 1)
	go func() {
		var (
			buf [65 * 1024]byte
		)
		n, err := r.Read(buf[:])
		if err != nil {
			close(out)
			return
		}
		out <- string(buf[:n])
	}()

	select {
	case str := <-out:
		return str
	case <-time.After(s.ReadWait):
		//log.Printf("Timed Out!\n")
		return ""
	}
}

// Close attempts to close the session and connection.
func (s *SSHRunner) Close() {
	s.sess.Close()
	s.conn.Close()
}

// Gather sends commands over SSH and returns a match to m (regex) defined in our collector
// Depends on google/expect which is not cross platform.
//func (s *SSHRunner) Gather(cmd []string, m *regexp.Regexp) string {
//	// Spawn expect
//	e, _, err := expect.SpawnSSH(s.conn, 10*time.Second)
//	if err != nil {
//		log.Fatalf("Failed to spawn expect: %v", err)
//	}
//	defer e.Close()
//
//	// Send each command
//	for c := range cmd {
//		err := e.Send(cmd[c] + "\n")
//		if err != nil {
//			log.Fatalf("Error sending command %v: %v", c, err)
//		}
//	}
//
//	// Wait for a match to "m" (regex passed from collector)
//	_, match, err := e.Expect(m, 10*time.Second)
//	if err != nil {
//		log.Fatalf("Error matching config: %v", err)
//	}
//
//	// return config block
//	return match[0]
//}
