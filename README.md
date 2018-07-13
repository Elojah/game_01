# GAME_01
> GAME_01 is a multi services scalable MMORPG server

GAME_01 is an UDP client/server with its own ACK protocol. Client writes user action and send them to API while receiving world data from sync. Core establish world rules and events order.
```
client <-> api -> core -> redis-lru
redis-lru -> sync -> client
```
Authentication and char creation/connect is handled by auth and associate a session token at each signin.
Revoker regularly revokes unused tokens.
GAME_01 also comes with a Tool API to create world data like entities/abilities/sectors.

## Code architecture
```
 _bin #executables
|
|_cmd_ #executables code
|     |_api #UDP API for game events + ACK client
|     |_auth #HTTPS JSON API for signin/connect
|     |_client #client/server to communicate with API and JSON serialize
|     |_core #order and apply game events
|     |_revoker #revoke unused tokens
|     |_sync #send entity data to clients
|     |_tool #HTTPS JSON API for world data. Must be private.
|
|_configs #config files
|
|_docker #docker files
|
|_pkg_ #common code
|     |_ability #domain
|     |_account #domain
|     |_dto #external format input/output
|     |_entity #domain
|     |_event #domain
|     |_geometry #domain
|     |_infra #domain
|     |_mocks #domain mocks
|     |_sector #domain
|     |_storage #domain mapper implementations
|     |_ulid #domain
|     |_usecase #domain aggregations
|
|_templates #example template files for tool
|
|_vendor #vendoring packages (dep)
```

## Installation

OS X & Linux & Windows:

```sh
go get -u github.com/elojah/game_01
```

## Development setup

```sh
# Start services
docker-compose -d
make dep
make sync && bin/game_sync configs/config_sync.json
make core && bin/game_core configs/config_core.json
make api && bin/game_api configs/config_api.json
make auth && bin/game_auth configs/config_auth.json
make tool && bin/game_tool configs/config_tool.json
```

## Usage example

```sh
# Fill static data
curl -k -X POST https://127.0.0.1:8081/entity/template -d @templates/entity_templates.json
curl -k -X POST https://127.0.0.1:8081/sector -d @templates/sector.json
curl -k -X POST https://127.0.0.1:8081/sector/starter -d @templates/sector_starter.json

# Obtain access token
curl -k -X POST https://127.0.0.1:8080/subscribe -d '{"username": "test", "password": "test"}'
curl -k -X POST https://127.0.0.1:8080/signin -d '{"username": "test", "password": "testtest"}'
{"ID":"01CJ9WDGGFQBNYWF514C6R6CQ6"}
curl -k -X POST https://127.0.0.1:8080/pc/create -d '{"token":"01CJ9WDGGFQBNYWF514C6R6CQ6","type":"01CE3J5ASXJSVC405QTES4M221", "name": "roger_lemour"}'
# Token is token obtained at signin and type is an entity ID described in templates/entity_templates.json.
curl -k -X POST https://127.0.0.1:8080/pc/list -d '{"token":"01CJ9WDGGFQBNYWF514C6R6CQ6"}'
[{"id":"01CJ9Y4S1PNSZ81NS8T3SR2DY6","type":"00000000000000000000000000","name":"mesmerist","hp":150,"mp":250,"position":{"Coord":{"x":39.19956060954395,"y":37.77876652333657,"z":36.315239570760646},"SectorID":"01CF001HTBA3CDR1ERJ6RF183A"}}]
curl -k -X POST https://127.0.0.1:8080/pc/connect -d '{"token":"01CJ9WDGGFQBNYWF514C6R6CQ6","target":"01CJ9Y4S1PNSZ81NS8T3SR2DY6"}'
{"ID":"01CJ9Y88P37R7A6KY83AP7ZP77"}
# Target is a PC ID in /list results

# Paste token in config_client.json: {... "app": {"token": 01CJ9WDGGFQBNYWF514C6R6CQ6,...}}
make client && bin/game_client configs/config_client.json
{"type":"move","action":{"source":"01CJ9Y88P37R7A6KY83AP7ZP77","target":"01CJ9Y88P37R7A6KY83AP7ZP77","position":{"X":94.0164,"Y":80.5287,"Z":70.7539}}}
...

# Disconnect PC only (may reconnect with same token)
curl -k -X POST https://127.0.0.1:8080/pc/disconnect -d '{"token": "01CJ9WDGGFQBNYWF514C6R6CQ6"}'
# Signout
curl -k -X POST https://127.0.0.1:8080/signout -d '{"username": "test", "token": "01CJ9WDGGFQBNYWF514C6R6CQ6"}'

```

_For more examples and usage, please refer to the [Wiki][wiki]._

## Release History

* 0.0.1
    * Achilles - Work in progress
* 0.0.2
    * Agni - Basic features

See [trello](https://trello.com/b/GX9gz3Js/game01) for more informations.

## Meta

Elojah – swingcastor@gmail.com

Distributed under the GNU AFFERO GENERAL PUBLIC license. See ``LICENSE`` for more information.

https://github.com/elojah/
