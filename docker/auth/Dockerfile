# builder
FROM golang:latest AS builder
WORKDIR /auth
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make auth

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /auth/bin/game_auth .
ENTRYPOINT ["/app/game_auth", "/app/auth.json"]
