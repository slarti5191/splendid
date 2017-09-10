package collectors

import "github.com/slarti5191/splendid"

func CollectCisco() (Cmds splendid.Commands) {
	// set commands to their expected output (last line)
	Cmds.Commands["set pager"] = ""
	Cmds.Commands["show run"] = "#"
	return Cmds
}