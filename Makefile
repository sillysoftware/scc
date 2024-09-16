all: build

build:
	@go build -o bin ./...

test:
	@go test ./...

clean:
	@go clean
