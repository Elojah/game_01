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
{"ID":"01CD05WMYFFMTHWCNE3PZNWPVK"}
// This token must be used as client token id
// Paste it in config_client.json: {... "app": {"token": 01CD05WMYFFMTHWCNE3PZNWPVK,...}}

> make client && bin/game_client bin/config_client.json
> {"type":"set_pc","action":{"type":"01CE3J5ASXJSVC405QTES4M221"}}
// FTM only way to retrieve generated pc is to check redis pc:token_id:pc_id
> {"type":"connect_pc","action":{"target":"01CGE6MJYQWSVNS5ZNNHHASKK9"}}
// FTM only way to retrieve generated entity is to check redis or added listener/sync
> {"type":"move","action":{"source":"01CGE7XCYKS6PH2925H4AP91MH","target":"01CGE7XCYKS6PH2925H4AP91MH","position":{"X":94.0164,"Y":80.5287,"Z":70.7539}}}
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
- [ ] Split auth (with set/list/connect PC) and remove all pc stuff elsewhere.
- [ ] Add server ack sending to client (and client resend)
- [ ] Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?* + first graphic client
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
