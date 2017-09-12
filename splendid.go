package splendid

import (
	"github.com/slarti5191/splendid/collectors"
)

const version = "0.0.0"

var usage = `splendid: Configuration trackoing for Network Devices.

Usage:
  splendid [options] <config>
  splendid -h --help
  splendid -v --version

Options:
  -w, --workspace <dir>     Specify workspace directory (default: ./splendid-workspace).
  -i, --interval <secs>     Collection interval in secs (default: 300).
  -c, --concurrency <num>   Concurrent device collections (default: 30).
  -t, --to <email@addr>     Send change notifications to this email.
  -f, --from <email@addr>   Send change notifications from this email.
  -s, --smtp <host:port>    SMTP server connection info (default: localhost:25).
  --insecure                Accept untrusted SSH device keys.
  --push                    Do a "git push" after committing changed configs.
  --timeout <secs>          Device collection timeout in secs (default: 60).
  --web                     Run an HTTP status server.
  --weblisten <host:port>   Host and port to use for HTTP status server (default: localhost:5000).
  -v --version                 Show version.
  -h, --help                Print this message.
`

func Init() {
	// Get global configs
	Conf := SetConfigs()
	// Set up DeviceConfig
	Dev := new(DeviceConfig)
	Dev.Method = "cisco"
	// Get Commands out of collectors.Generate
	Cmds := collectors.Generate(string(Dev.Method))
	// Kick off a collector
	RunCollector(Dev, Conf, Cmds)
}
