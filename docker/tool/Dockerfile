# builder
FROM golang:latest AS builder
WORKDIR /tool
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make tool

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /tool/bin/game_tool .
ENTRYPOINT ["/app/game_tool", "/app/tool.json"]
