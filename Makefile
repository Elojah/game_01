PACKAGE   = game
DATE     ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

ifneq ($(wildcard /snap/go/current/bin/go),)
	GO = /snap/go/current/bin/go
else ifneq ($(shell which go1.11),)
	GO = go1.11
else
	GO = go
endif

ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

GODOC       = godoc
GOFMT       = gofmt

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

# Help
go-version:
	$Q echo $(GO)

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
.PHONY: proto
proto:
	$(info $(M) running protobuf…) @
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. component.proto
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. ability.proto
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. effect_feedback.proto
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. feedback.proto
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. targets.proto
	$Q cd pkg/ability  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. starter.proto
	$Q cd pkg/account  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. account.proto
	$Q cd pkg/account  && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. token.proto
	$Q cd pkg/geometry && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. position.proto
	$Q cd pkg/item     && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. item.proto
	$Q cd pkg/entity   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. dto.proto
	$Q cd pkg/entity   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. entity.proto
	$Q cd pkg/entity   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. inventory.proto
	$Q cd pkg/entity   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. spawn.proto
	$Q cd pkg/event    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. action.proto
	$Q cd pkg/event    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. dto.proto
	$Q cd pkg/event    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. event.proto
	$Q cd pkg/event    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. trigger.proto
	$Q cd pkg/infra    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. ack.proto
	$Q cd pkg/infra    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. q_action.proto
	$Q cd pkg/infra    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. sequencer.proto
	$Q cd pkg/infra    && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. recurrer.proto
	$Q cd pkg/sector   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. entities.proto
	$Q cd pkg/sector   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. items.proto
	$Q cd pkg/sector   && protoc -I=. -I=$(GOPATH)/src --gogoslick_out=. sector.proto

# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor

# Tidy
.PHONY: tidy
tidy:
	$(info $(M) running go mod tidy…) @
	$Q $(GO) mod tidy

# Check
.PHONY: check
check: vendor lint test

# Lint
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run

# Test
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

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
