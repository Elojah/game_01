# builder
FROM golang:latest AS builder
WORKDIR /api
COPY . .
RUN make tidy
RUN make api

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /api/bin/game_api /app/
COPY --from=builder /api/config/api.json /app/
ENTRYPOINT ["game_api", "api.json"]
