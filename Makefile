all: build

build:
	@-mkdir bin
	@go build -o bin ./...

test:
	@go test ./...

clean:
	@go clean
	@-rm -rf bin/
