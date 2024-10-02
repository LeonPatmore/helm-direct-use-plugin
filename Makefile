setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

run:
	go run cmd/directuse/main.go

build:
	go build cmd/directuse/main.go

lint:
	golangci-lint run --timeout=3m --verbose

format:
	gofmt -s -w .

test:
	go test -v ./...
