package languages

var ProgrammingLanguages = map[string]Runner{
	"JavaScript": JavaScriptRunner{
		LanguageConfig{
			DockerImage:  "node:alpine",
			Extension:    "js",
			TestsSupport: true,
		},
	},

	"TypeScript": TypeScriptRunner{
		LanguageConfig{
			DockerImage:  "oven/bun:alpine",
			Extension:    "ts",
			TestsSupport: true,
		},
	},

	"Go": GolangRunner{
		LanguageConfig{
			DockerImage:  "golang:alpine",
			Extension:    "go",
			TestsSupport: true,
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
			DockerImage:  "rust:alpine",
			Extension:    "rs",
			TestsSupport: false,
		},
	},

	"PHP": PhpRunner{
		LanguageConfig{
			DockerImage:  "php:alpine",
			Extension:    "php",
			TestsSupport: false,
		},
	},

	"Bash": BashRunner{
		LanguageConfig{
			DockerImage:  "alpine",
			Extension:    "sh",
			TestsSupport: false,
		},
	},

	"C++": CppRunner{
		LanguageConfig{
			DockerImage:  "gcc",
			Extension:    "cpp",
			TestsSupport: false,
		},
	},

	"C": CRunner{
		LanguageConfig{
			DockerImage:  "gcc",
			Extension:    "c",
			TestsSupport: false,
		},
	},
}
