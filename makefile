.DEFAULT_GOAL := build

fmt:
	go fmt .

lint: fmt
	golint .

vet: fmt
	go vet .

build: vet
	go build -o my-cloud-home-go

run: build
	go run main.go
