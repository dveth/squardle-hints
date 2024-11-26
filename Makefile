.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet:
	go vet ./...
build: vet
	go build -o squardle-hints.exe ./cmd/squardle-hints/ 
clean: 
	go clean
test: 
	go test ./tests/
