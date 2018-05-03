PACKAGE   = game
DATE     ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell git describe --tags --always --dirty --match = v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

GO        = go
GODOC     = godoc
GOFMT     = gofmt
GOLINT    = gometalinter

API       = api
CLIENT    = client
AUTH      = auth
PLAYER    = player

FBSDIR    = .

V         = 0
Q         = $(if $(filter 1,$V),,@)
M         = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all

all: api client auth

# Executables
api:
	$(info $(M) building executable api…) @ ## Build program binary
	$Q cd cmd/$(API) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

client:
	$(info $(M) building executable client…) @ ## Build program binary
	$Q cd cmd/$(CLIENT) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(CLIENT)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(CLIENT)_$(VERSION) bin/$(PACKAGE)_$(CLIENT)

auth:
	$(info $(M) building executable auth…) @ ## Build program binary
	$Q cd cmd/$(AUTH) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(AUTH)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(AUTH)_$(VERSION) bin/$(PACKAGE)_$(AUTH)

player:
	$(info $(M) building executable player…) @ ## Build program binary
	$Q cd cmd/$(PLAYER) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(PLAYER)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(PLAYER)_$(VERSION) bin/$(PACKAGE)_$(PLAYER)

.PHONY: gen
gen:
	$(info $(M) running gencode…) @
	$Q cd dto && gencode go -schema=message.schema -package dto
	$Q cd storage &&\
		gencode go -schema=token.schema -package storage &&\
		gencode go -schema=event.schema -package storage &&\
		gencode go -schema=account.schema -package storage

# Dependencies
.PHONY: dep
dep:
	$(info $(M) building vendor…) @
	$Q dep ensure

# Check
.PHONY: check
check: lint test

# Tests
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Tools
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…) @
	$Q GO_VENDOR=1 $(GOLINT) "--vendor" \
					"--disable=gotype" \
					"--disable=vetshadow" \
					"--disable=gocyclo" \
					"--fast" \
					"--json" \
					"./..." \
			| grep -v schema.gen.go

.PHONY: fmt
fmt:
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc:
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: clean
clean:
	$(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin/$(PACKAGE)_$(API)*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
