package collectors

// Generate initializes Commands based on DeviceConfig.Method
func Generate(m string) (Cmds Commands) {
	switch m {
	case "cisco":
		return CiscoCmd()
	default:
		return CiscoCmd()
	}
}
