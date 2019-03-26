# GAME_01

**WORK IN PROGRESS**

[![CircleCI](https://circleci.com/gh/Elojah/game_01/tree/master.svg?style=svg)](https://circleci.com/gh/Elojah/game_01/tree/master)

> GAME_01 is a multi services scalable MMORPG server

GAME_01 is a set of services and tools to build a MMORPG server the easiest way possible
*UE4 integration in progress*

## Installation

OS X & Linux & Windows:

```sh
go get -u github.com/elojah/game_01
```

## Development setup

```sh
# Start services
docker-compose -d
make vendor
make sync && bin/game_sync configs/config_sync.json
make core && bin/game_core configs/config_core.json
make api && bin/game_api configs/config_api.json
make auth && bin/game_auth configs/config_auth.json
make tool && bin/game_tool configs/config_tool.json
```

## Usage example
*TODO*
_For more examples and usage, please refer to the [Wiki][wiki]._

## Release History

* 0.0.1
    * Achilles - Work in progress
* 0.0.2
    * Agni - Basic features

See [trello](https://trello.com/b/GX9gz3Js/game01) for more informations.

## How it works
```
client <-> api -> core -> redis-lru
redis-lru -> sync -> client
```
Authentication and char creation/connection is handled by auth and associate a session token at each signin.
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
|     |_ability_ # domain
|     |         |_svc # service/usecases
|     |         |_srg # storage/database
|     |
|     |_account
|     |_entity
|     |_event
|     |_geometry
|     |_infra
|     |_sector
|     |_ulid
|
|_static #example template files for tool
|
|_vendor #vendoring packages (dep)
```

## Meta

Elojah – swingcastor@gmail.com

Distributed under the GNU AFFERO GENERAL PUBLIC license. See ``LICENSE`` for more information.

https://github.com/elojah/
