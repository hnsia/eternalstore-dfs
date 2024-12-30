build:
	@go build -o bin/es

run: build
	@./bin/es

test:
	@go test ./... -v