FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o lisp ./...

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/lisp .

CMD ["./lisp"]
