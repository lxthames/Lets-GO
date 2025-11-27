DEFAULT_GOAL := build
.PHONY: fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o hello_world.exe

clean:
	go clean
	del /S /Q *.exe 2>nul || exit 0
