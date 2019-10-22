#!/bin/sh
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/documents/openapi.yaml -g go-gin-server -o /local/generate/gin-server
