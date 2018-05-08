# game_01

## How to start

```
> docker-compose -d
> make dep
> make auth && bin/game_auth bin/config_auth.json
> make api && bin/game_api bin/config_api.json
> make sequencer && bin/game_sequencer bin/config_sequencer.json
> curl -k -X POST https://127.0.0.1:8080/subscribe -d '{"username": "test", "password": "test"}'
> curl -k -X POST https://127.0.0.1:8080/login -d '{"username": "test", "password": "test"}'
{"ID":"01CD05WMYFFMTHWCNE3PZNWPVK"}
// This token must be used for next client
> make client && bin/game_client bin/config_client.json
```

## TODO
- [x] Remove NATS streaming
- [x] Change log to zap (uber faster log)
- [x] Add context + use it with sync.errgroup
- [x] Use TCP + rename UDP to mux
- [x] Create new https service to create new token
- [x] Create new bin to read events
- [x] Add NATS mqueue to cancel replay mechanism (in context usage ?)
- [ ] Fix NATS slow consumer
- [ ] Add server ack sending to client (and client resend)
- [ ] ack service
- [ ] Think about entity interactions limit to "what's around" to scale efficiently
- [ ] handler controller
    + resolve TODOs
- [ ] `tile38` Entity Service
- [ ] `storage/entity.go` to group use_cases
- [ ] Response server to update all clients with delta compression named `sync`
- [ ] Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?*
