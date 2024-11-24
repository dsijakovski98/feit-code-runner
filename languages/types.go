package languages

type LanguageConfig struct {
	DockerImage string
	Extension   string
}

type Runner interface {
	// filePath is WITH file extension
	CommandChain(filePath string) []string
	ExtraCommands(filePath string, containerId string) (string, error)
	GetConfig() LanguageConfig
}
