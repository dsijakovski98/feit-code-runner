package config

import (
	"fmt"
	"os"
)

func GetAllowedOrigins() []string {
	fmt.Println(os.Getenv("FEIT_CODE_ENV"))
	prod := os.Getenv("FEIT_CODE_ENV") == "production"

	if prod {
		return []string{"https://app.feitcode.com"}
	}

	return []string{"http://localhost:3001"}
}
