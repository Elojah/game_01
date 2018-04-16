# game_01

## How to start

```
docker-compose -d
make api && bin/game_api bin/config_api.json
make client && bin/game_client bin/config_client.json
```

## TODO
- HTTPS service for users with token creation.
- Response server to update all clients with delta compression.
