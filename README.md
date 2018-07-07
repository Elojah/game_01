# GAME_01
> GAME_01 is a multi services scalable MMORPG server

GAME_01 is an UDP client/server with its own ACK protocol. Client writes user action and send them to API while receiving world data from sync. Core establish world rules and events order.
```
client <-> api -> core -> redis-lru
redis-lru -> sync -> client
```
Authentication and char creation/connect is handled by auth and associate a session token at each login.
Revoker regularly revokes unused tokens.
GAME_01 also comes with a Tool API to create world data like entities/abilities/sectors.

## Installation

OS X & Linux & Windows:

```sh
go get -u github.com/elojah/game_01
```

## Usage example

```sh
// Fill static data
curl -k -X POST https://127.0.0.1:8081/entity/template -d @templates/entity_templates.json
curl -k -X POST https://127.0.0.1:8081/sector -d @templates/sector.json
curl -k -X POST https://127.0.0.1:8081/sector/starter -d @templates/sector_starter.json

// Obtain access token
curl -k -X POST https://127.0.0.1:8080/subscribe -d '{"username": "test", "password": "test"}'
curl -k -X POST https://127.0.0.1:8080/login -d '{"username": "test", "password": "test"}'
{"ID":"01CHNKE1NZERJ37PHVAQ0STBWQ"}
curl -k -X POST https://127.0.0.1:8080/pc/create -d '{"token":"01CHNKE1NZERJ37PHVAQ0STBWQ","type":"01CE3J5ASXJSVC405QTES4M221"}'
// Token is token obtained at login and type is an entity ID described in templates/entity_templates.json.
curl -k -X POST https://127.0.0.1:8080/pc/list -d '{"token":"01CHNKE1NZERJ37PHVAQ0STBWQ"}'
[{"id":"01CHNKEKCWT34W50C2JY18YJYE","type":"00000000000000000000000000","name":"mesmerist","hp":150,"mp":250,"position":{"Coord":{"x":39.19956060954395,"y":37.77876652333657,"z":36.315239570760646},"SectorID":"01CF001HTBA3CDR1ERJ6RF183A"}}]
curl -k -X POST https://127.0.0.1:8080/pc/connect -d '{"token":"01CHNKE1NZERJ37PHVAQ0STBWQ","target":"01CHNKEKCWT34W50C2JY18YJYE"}'
{"ID":"01CHNKFCTJWZ4PV3BJGFZ1QY4R"}
// Target is a PC ID in /list results

// Paste token in config_client.json: {... "app": {"token": 01CHNKE1NZERJ37PHVAQ0STBWQ,...}}
make client && bin/game_client bin/config_client.json
{"type":"move","action":{"source":"01CHNKFCTJWZ4PV3BJGFZ1QY4R","target":"01CHNKFCTJWZ4PV3BJGFZ1QY4R","position":{"X":94.0164,"Y":80.5287,"Z":70.7539}}}
```

_For more examples and usage, please refer to the [Wiki][wiki]._

## Development setup

```sh
// Start services
docker-compose -d
make dep
make sync && bin/game_sync bin/config_sync.json
make core && bin/game_core bin/config_core.json
make api && bin/game_api bin/config_api.json
make auth && bin/game_auth bin/config_auth.json
make tool && bin/game_tool bin/config_tool.json
```

## Release History

* 0.0.1
    * Work in progress

## Meta

Elojah – swingcastor@gmail.com

Distributed under the GNU AFFERO GENERAL PUBLIC license. See ``LICENSE`` for more information.

https://github.com/elojah/

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[travis-image]: __
[travis-url]: __
[wiki]: https://github.com/yourname/yourproject/wiki

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
- [x] Add up/down mechanic
- [x] Save pool ID in token.
- [x] Save PC + Delete all token-associated data.
- [x] Add use cases for entity (create/delete) and token(create/delete)
- [x] Switch nats to redis pub/sub
- [x] Add Pool for listeners + ListenerMapper Set/Get
- [x] Del for everything (not yet but token/entity ok)
- [x] Put everything as usecase and use (almost) only them in controllers (not yet but entity/token ok)
- [ ] Integration test binary with correct set
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
