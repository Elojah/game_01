# builder
FROM golang:latest AS builder
WORKDIR /integration
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make integration

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /integration/bin/game_integration .
ENTRYPOINT ["/app/game_integration"]
