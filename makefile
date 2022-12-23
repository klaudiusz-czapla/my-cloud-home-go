.DEFAULT_GOAL := build

fmt:
	goimports -l -w . && go fmt .

test:
	go clean -testcache && go test ./...

lint: fmt
	golangci-lint run .

vet: fmt
	go vet .

build: vet
	go build -o my-cloud-home-go

run: build
	go run main.go
