FROM danielgtaylor/apisprout:latest

COPY ./documents/openapi.yaml ./openapi.yaml

ENV SWAGGER_JSON=/openapi.yaml

EXPOSE 8000

ENTRYPOINT /usr/local/bin/apisprout /openapi.yaml --watch
