all: build

build: clean
	@-mkdir bin
	@go build -o bin/ ./...
	@mv bin/go bin/lisp

test:
	@go test ./...

clean:
	@go clean
	@-rm -rf bin/
