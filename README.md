# game_01

## How to start

```
docker-compose -d
make dep
make api && bin/game_api bin/config_api.json
make client && bin/game_client bin/config_client.json
```

## TODO
- ip service (+ token association)
- ack service
- handler controller
    + resolve TODOs
- `tile38` Actor Service
- `storage/actor.go` to group use_cases
- HTTPS service for users with token creation based on PG named `auth`
- Response server to update all clients with delta compression named `sync`
- Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?*
