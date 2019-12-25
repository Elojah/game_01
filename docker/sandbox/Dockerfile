# builder
FROM golang:latest AS builder
WORKDIR /sandbox
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make sandbox

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /sandbox/bin/game_sandbox .
ENTRYPOINT ["/app/game_sandbox", "/app/sandbox.json"]
