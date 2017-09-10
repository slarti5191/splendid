package splendid


type Commands struct {
	Commands map[string]string
}
// RunCollector collects configs
// Grab global configs as Conf, device specific commands as Cmd
// Cmd should be set up under collectors/
func RunCollector(Conf *SplendidConfig, Cmd *Commands) {
	// iterate over commands, expect matching output, fail otherwise

}