# builder
FROM golang:latest AS builder
WORKDIR /revoker
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make revoker

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /revoker/bin/game_revoker .
ENTRYPOINT ["/app/game_revoker", "/app/revoker.json"]
