# game_01

## How to start

```
// Start services
> docker-compose -d
> make dep
> make sync && bin/game_sync bin/config_sync.json
> make core && bin/game_core bin/config_core.json
> make api && bin/game_api bin/config_api.json
> make auth && bin/game_auth bin/config_auth.json
> make tool && bin/game_tool bin/config_tool.json

// Fill static data
> curl -k -X POST https://127.0.0.1:8081/entity/template -d @templates/entity_templates.json
> curl -k -X POST https://127.0.0.1:8081/sector -d @templates/sector.json

// Obtain access token
> curl -k -X POST https://127.0.0.1:8080/subscribe -d '{"username": "test", "password": "test"}'
> curl -k -X POST https://127.0.0.1:8080/login -d '{"username": "test", "password": "test"}'
{"ID":"01CGH6TXS6PAXHWY79KCGKVZGV"}
> curl -k -X POST https://127.0.0.1:8080/pc/create -d '{"token":"01CGH6TXS6PAXHWY79KCGKVZGV","type":"01CE3J5ASXJSVC405QTES4M221"}'
// Token is token obtained at login and type is an entity ID described in templates/entity_templates.json.
> curl -k -X POST https://127.0.0.1:8080/pc/list -d '{"token":"01CGH6TXS6PAXHWY79KCGKVZGV"}'
[{"id":"01CGH6VKAKN4BQRQWF5NX6ZNS2","type":"00000000000000000000000000","name":"mesmerist","hp":150,"mp":250,"position":{"Coord":{"x":39.19956060954395,"y":37.77876652333657,"z":36.315239570760646},"SectorID":"01CF001HTBA3CDR1ERJ6RF183A"}}]
> curl -k -X POST https://127.0.0.1:8080/pc/connect -d '{"token":"01CGH6TXS6PAXHWY79KCGKVZGV","target":"01CGH6VKAKN4BQRQWF5NX6ZNS2"}'
{"ID":"01CGH399MZYQZX71V36YH4XZEW"}
// Target is a PC ID in /list results

// Paste token in config_client.json: {... "app": {"token": 01CGH6TXS6PAXHWY79KCGKVZGV,...}}
> make client && bin/game_client bin/config_client.json
> {"type":"move","action":{"source":"01CGH6VKAKN4BQRQWF5NX6ZNS2","target":"01CGH6VKAKN4BQRQWF5NX6ZNS2","position":{"X":94.0164,"Y":80.5287,"Z":70.7539}}}
```

## TODO
- [x] Remove NATS streaming
- [x] Change log to zap (uber faster log)
- [x] Add context + use it with sync.errgroup
- [x] Use TCP + rename UDP to mux
- [x] Set new https service to create new token
- [x] Set new bin to read events
- [x] Add NATS mqueue to cancel replay mechanism (in context usage ?)
- [x] Fix NATS slow consumer, mouais
- [x] Add `sequencer_test.go`, 100% plz (ok)
- [x] Add state/entity service impl + interactions
- [x] Handle token permissions/entity actions (linked to above)
- [x] Refacto skill mechanic to be like class definition (json template + tool)
- [x] Cast/Skill can have multiple effects on multiple targets
- [x] AbilityFeedback service get/set
- [x] Add core app handler skill
- [x] Implement SkillFeedback mechanism
- [x] Response server to update all clients with delta compression named `sync`
    + [x] Think about entity interactions limit to "what's around" to scale efficiently (`tile38` Entity Service) ?
    + [x] recurrer test 100%
- [x] ack service
- [x] Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?* + first graphic client
- [x] Split auth (with set/list/connect PC) and remove all pc stuff elsewhere.
- [x] Add token revoker after X incativity
- [x] Change EntitySubset key string into ID ID
- [x] Add redis services for core and syncs ids apps
- [x] Add redis service for initial positions.
- [x] Move all Start into Dial (for clean/up down)
- [ ] Add up/down mechanic
- [ ] Save pool ID in listeners.
- [ ] Add use cases for entity (create/delete) and token(create/delete)
- [ ] Add server ack sending to client (and client resend)
- [ ] Add minimal graphic interface to client (minimal calls and print entity states)
- [ ] Implement tool to generate/check/visualize sectors and entity movements
- [ ] Add context everywhere

## TO DEBATE
- [ ] Add sequencer cancelling mechanism (CancelEvent(id)->id(don't do me)->event...)
- [ ] Add cast time mechanic

## RANDOM
- [ ] `Trickster` entity can switch position with his own entities and switch them
- [ ] `Mesmerist` entity can take control of enemy entities
- [ ] `Inquisitor` entity can merge entities (allies/enemies)
- [ ] `Totemist` entity can clone his own entities
- [ ] `Scavenger` entity can sacrify his own entities
