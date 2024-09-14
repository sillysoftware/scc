FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN make build

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/lisp .

CMD ["./bin/lisp"]
