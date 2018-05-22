# game_01

## How to start

```
> docker-compose -d
> make dep
> make auth && bin/game_auth bin/config_auth.json
> make api && bin/game_api bin/config_api.json
> make core && bin/game_core bin/config_core.json
> curl -k -X POST https://127.0.0.1:8080/subscribe -d '{"username": "test", "password": "test"}'
> curl -k -X POST https://127.0.0.1:8080/login -d '{"username": "test", "password": "test"}'
{"ID":"01CD05WMYFFMTHWCNE3PZNWPVK"}
// This token must be used as client token id
> make client && bin/game_client bin/config_client.json
```

##### Tool
```
> make add-templates
> make show-templates
> make add-skills
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
- [ ] Cast/Skill can have multiple effects on multiple targets
- [ ] Add core app handler skill
- [ ] Implement SkillFeedback mechanism
- [ ] Add sequencer cancelling mechanism (CancelEvent(id)->id(don't do me)->event...)
- [ ] Response server to update all clients with delta compression named `sync`
- [ ] Add server ack sending to client (and client resend)
- [ ] ack service
- [ ] Think about entity interactions limit to "what's around" to scale efficiently (`tile38` Entity Service) ?
- [ ] Edit `client` to make it sensitive to `sync` calls and save in a local *rocksdb ?* + first graphic client
- [ ] Add context everywhere

## RANDOM
- [ ] `Trickster` entity can switch position with his own entities and switch them
- [ ] `Mesmerist` entity can take control of enemy entities
- [ ] `Inquisitor` entity can merge entities (allies/enemies)
- [ ] `Totemist` entity can clone his own entities
- [ ] `Scavenger` entity can sacrify his own entities
