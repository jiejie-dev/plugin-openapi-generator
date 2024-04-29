FROM golang:1.22 as builder

WORKDIR /app

COPY ./go.* .
RUN go mod download

COPY . .
RUN go build -o ./plugin

FROM openapitools/openapi-generator-cli as runner
COPY --from=builder /app/plugin ./plugin

# ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

ENTRYPOINT [ "./plugin" ]