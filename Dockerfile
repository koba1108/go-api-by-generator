#FROM golang:latest

#WORKDIR /go/src/github/go-api-by-generator
#ENV GO111MODULE=on

#EXPOSE 8080

FROM swaggerapi/swagger-ui

ADD . /documents/openapi.yaml:/openapi.yaml
ENV SWAGGER_JSON=/openapi.yaml

EXPOSE 8080

# CMD ["go", "run", "main.go"]
