# game_01

## How to start

```
docker-compose -d
make dep
make api && bin/game_api bin/config_api.json
make client && bin/game_client bin/config_client.json
```

## TODO
- `tile38` Actor Service
- `storage/actor.go` to group use_cases
- `cmd/api/actor.go` controllers
  + token verification
  + attack in 1 area/move
  + ack state save
- HTTPS service for users with token creation based on PG named `auth`
- Response server to update all clients with delta compression named `sync`
