.DEFAULT_GOAL := build
.PHONY: vet build fmt
fmt : 
		go fmt ./...
vet: fmt
		go vet ./...
clean: vet
		go clean -i -r -cache
build: clean
		go build -o dummy