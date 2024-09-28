setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

run:
	go run cmd/example/example.go

build:
	go build cmd/example/example.go

lint:
	golangci-lint run --timeout=3m

format:
	gofmt -s -w .

test:
	go test -v ./...

test-plugin:
	helm plugin uninstall direct-use
	helm plugin install .
