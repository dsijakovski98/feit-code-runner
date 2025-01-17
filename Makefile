run:
	go run main.go

build:
	go build -o bin/feit-code-runner

start:
	./bin/feit-code-runner

setup:
	cd parsers/js-ts && bun install
	cd parsers/go && go install
