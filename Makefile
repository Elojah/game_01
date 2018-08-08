PACKAGE   = game
DATE     ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

GO        = go
GODOC     = godoc
GOFMT     = gofmt
GOLINT    = gometalinter

API         = api
CLIENT      = client
AUTH        = auth
CORE        = core
SYNC        = sync
TOOL        = tool
REVOKER     = revoker
INTEGRATION = integration

V         = 0
Q         = $(if $(filter 1,$V),,@)
M         = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all

all: client auth api core sync revoker tool

# Executables
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

api:
	$(info $(M) building executable api…) @ ## Build program binary
	$Q cd cmd/$(API) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

core:
	$(info $(M) building executable core…) @ ## Build program binary
	$Q cd cmd/$(CORE) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(CORE)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(CORE)_$(VERSION) bin/$(PACKAGE)_$(CORE)

sync:
	$(info $(M) building executable sync…) @ ## Build program binary
	$Q cd cmd/$(SYNC) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(SYNC)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(SYNC)_$(VERSION) bin/$(PACKAGE)_$(SYNC)

revoker:
	$(info $(M) building executable revoker…) @ ## Build program binary
	$Q cd cmd/$(REVOKER) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(REVOKER)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(REVOKER)_$(VERSION) bin/$(PACKAGE)_$(REVOKER)

tool:
	$(info $(M) building executable tool…) @ ## Build program binary
	$Q cd cmd/$(TOOL) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(TOOL)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(TOOL)_$(VERSION) bin/$(PACKAGE)_$(TOOL)

integration:
	$(info $(M) building executable integration…) @ ## Build program binary
	$Q cd cmd/$(INTEGRATION) &&  $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o ../../bin/$(PACKAGE)_$(INTEGRATION)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(INTEGRATION)_$(VERSION) bin/$(PACKAGE)_$(INTEGRATION)

# Utils
.PHONY: gen
gen:
	$(info $(M) running gencode…) @
	$Q gencode go -unsafe -def-types=false -package=ability -schema=schemas/ability.schema -out=pkg/ability/ability.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=account -schema=schemas/account.schema -out=pkg/account/account.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=infra -schema=schemas/ack.schema -out=pkg/infra/ack.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=event -schema=schemas/action.schema -out=pkg/event/action.schema.gen.go -ignore=Position,Vec3
	$Q gencode go -unsafe -def-types=false -package=event -schema=schemas/dto.schema -out=pkg/event/dto.schema.gen.go -ignore=Vec3,Position,Move,Cast
	$Q gencode go -unsafe -def-types=false -package=entity -schema=schemas/entity.schema -out=pkg/entity/entity.schema.gen.go -ignore=Position,Vec3
	$Q gencode go -unsafe -def-types=false -package=event -schema=schemas/event.schema -out=pkg/event/event.schema.gen.go -ignore=Vec3,Position,Move,Cast,Feedback
	$Q gencode go -unsafe -def-types=false -package=infra -schema=schemas/listener.schema -out=pkg/infra/listener.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=infra -schema=schemas/recurrer.schema -out=pkg/infra/recurrer.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=sector -schema=schemas/sector.schema -out=pkg/sector/sector.schema.gen.go -ignore=Vec3,Position
	$Q gencode go -unsafe -def-types=false -package=sector -schema=schemas/sector_entities.schema -out=pkg/sector/entities.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=account -schema=schemas/token.schema -out=pkg/account/token.schema.gen.go
	$Q gencode go -unsafe -def-types=false -package=geometry -schema=schemas/position.schema -out=pkg/geometry/position.schema.gen.go

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
					"--disable=goconst" \
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
	@rm -rf bin/$(PACKAGE)_*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
