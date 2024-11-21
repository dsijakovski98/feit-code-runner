package languages

type LanguageConfig struct {
	DockerImage string
	Extension   string
}

type Runner interface {
	// filePath is WITH file extension
	CommandChain(filePath string) []string
	ExtraCommands(containerId string) error
	GetConfig() LanguageConfig
}
