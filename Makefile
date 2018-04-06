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

FBSDIR    = .

V         = 0
Q         = $(if $(filter 1,$V),,@)
M         = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all

all: api

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

# Dependencies
.PHONY: dep
dep:
	$(info $(M) building vendor…) @
	$Q dep ensure -update

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
			| grep -v _generated.go
.PHONY: fmt
fmt:
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc:
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: fbs
fbs:
	$(info $(M) generating models…)	@ ## Generate pb.go files
	$(GOPATH)/src/github.com/google/flatbuffers/flatc\
		--go\
		--gen-onefile\
		--go-namespace game\
		-o $(FBSDIR)\
		$(FBSDIR)/*.fbs

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
