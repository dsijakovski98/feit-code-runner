package config

import (
	"os"
)

func GetEnv() string {
	env := os.Getenv("FEIT_CODE_ENV")

	if env == "" {
		env = "development"
	}

	return env
}

func GetAllowedOrigins() []string {
	env := GetEnv()
	prod := env == "production"

	if prod {
		return []string{"https://app.feitcode.com"}
	}

	return []string{"http://localhost:3001"}
}
