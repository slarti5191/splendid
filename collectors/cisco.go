package splendid

type Commands struct {
	Commands map[string]string
}

func CiscoCmd() (Cmds *Commands) {
	// set commands to their expected output (last line)
	Cmds.Commands["set pager"] = ""
	Cmds.Commands["show run"] = "#"
	return Cmds
}