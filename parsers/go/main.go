package main

import (
	"flag"
	"log"
	"os"

	"github.com/dsijakovski98/feit-code-runner/go-parser/utils"
)

var fileFlag = flag.String("file", "_", "Path to the file being parsed")

func ProcessCode(code string) (string, error) {
	cleanCode, err := utils.CleanupDebugs(code)
	if err != nil {
		return "", err
	}

	finalCode, err := utils.AppendPlaceholder(cleanCode)
	if err != nil {
		return "", err
	}

	return finalCode, nil
}

func main() {
	flag.Parse()

	bytes, err := os.ReadFile(*fileFlag)
	if err != nil {
		log.Fatal(err.Error())
	}

	parsedCode, err := ProcessCode(string(bytes))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := os.WriteFile(*fileFlag, []byte(parsedCode), 0644); err != nil {
		log.Fatal(err.Error())
	}
}
