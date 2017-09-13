package collectors

type Commands struct {
	Commands map[string]string
}

func CiscoCmd() (Cmds Commands) {
	// set commands to their expected output (last line)
	return Commands {
		map[string]string{
			"set pager": "",
			"show run":  "#",
		},
	}
}