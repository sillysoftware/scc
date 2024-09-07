all: run

run:
	@go run ./... demo/exit.lisp

test:
	@go test lisp/go/asm/asm_test.go
