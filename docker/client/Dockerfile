# builder
FROM golang:latest AS builder
WORKDIR /client
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make client

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /client/bin/game_client .
ENTRYPOINT ["/app/game_client", "/app/client.json"]
