package languages

type LanguageConfig struct {
	DockerImage  string
	Extension    string
	TestsSupport bool
}

type Runner interface {
	// filePath is WITH file extension
	GetConfig() LanguageConfig

	RunCommand(filePath string) []string
	ExtraRunCommands(filePath string, containerId string) (string, error)

	ParserDir() string
	ParserCommand(filePath string) (string, []string)
}
