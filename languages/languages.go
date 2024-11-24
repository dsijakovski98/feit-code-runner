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
			DockerImage: "oven/bun:alpine",
			Extension:   "ts",
		},
	},

	"Go": GolangRunner{
		LanguageConfig{
			DockerImage: "golang:alpine",
			Extension:   "go",
		},
	},

	"Python": PythonRunner{
		LanguageConfig{
			DockerImage: "python:alpine",
			Extension:   "py",
		},
	},

	"Rust": RustRunner{
		LanguageConfig{
			DockerImage: "rust:alpine",
			Extension:   "rs",
		},
	},
}
