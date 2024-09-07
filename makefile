all: build

build:
	@go build -o lispc ./...

test:
	@go test lisp/go/asm/asm_test.go

clean:
	@go clean
