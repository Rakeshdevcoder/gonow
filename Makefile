build:
	@go build -o bin/gobank ./cmd

run: build
	@./bin/gobank

test:
	@go test -v ./...