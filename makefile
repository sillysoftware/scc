all: run

run:
	@go run ./... demo/exit.lisp

test:
	@go test ./...
