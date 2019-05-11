# builder
FROM golang:latest AS builder
WORKDIR /api
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make api

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /api/bin/game_api .
ENTRYPOINT ["/app/game_api", "/app/api.json"]
