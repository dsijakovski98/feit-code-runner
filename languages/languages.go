package languages

var ProgrammingLanguages = map[string]Runner{
	"JavaScript": JavaScriptRunner{
		LanguageConfig{
			DockerImage: "node:alpine",
			Extension:   "js",
		},
	},

	"TypeScript": TypeScriptRunner{
		LanguageConfig{
			DockerImage: "oven/bun:latest",
			Extension:   "ts",
		},
	},
}
