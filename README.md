# game_01

## How to start

```
docker-compose -d
make dep
make api && bin/game_api bin/config_api.json
make client && bin/game_client bin/config_client.json
```

## TODO
- [x] Remove NATS streaming
- [x] Change log to zap (uber faster log)
- [x] Add context + use it with sync.errgroup
- [x] Use TCP + rename UDP to mux
- [ ] Add influxDB dep + impl entity/state service
- [ ] Add NATS mqueue to cancel replay mechanism (in context usage ?)
- [ ] Add server ack sending to client (and client resend)
- [ ] ack service
- [ ] Think about entity interactions limit to "what's around" to scale efficiently
- [ ] handler controller
    + resolve TODOs
- [ ] `tile38` Entity Service
- [ ] `storage/entity.go` to group use_cases
- [ ] HTTPS service for users with token creation based on PG named `auth`
- [ ] Response server to update all clients with delta compression named `sync`
- [ ] Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?*
