all: build

build:
	@go run ./... demo/exit.lisp

test:
	@go test lisp/go/asm/asm_test.go

clean:
	@go clean
