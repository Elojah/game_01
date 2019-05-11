# builder
FROM golang:latest AS builder
WORKDIR /sync
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make sync

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /sync/bin/game_sync .
ENTRYPOINT ["/app/game_sync", "/app/sync.json"]
