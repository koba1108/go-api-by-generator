FROM golang:latest

WORKDIR /go/src/go-api-by-generater
ENV GO111MODULE=on

EXPOSE 8080

CMD ["go", "run", "/go/src/main.go"]
