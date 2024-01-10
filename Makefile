#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
DOCKER := $(shell which docker)

build:
	go mod tidy \
	&& go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o colamosd cmd/colamosd/main.go
.PHONY: build
BINS+=colamosd

install:build
	cp -f colamosd ${GOPATH}/bin && colamosd version

init:
	ignite chain init --skip-proto -y

ignite-build:
	ignite chain build -y --debug --clear-cache --check-dependencies -v

# legacy version 0.11.6
protoVer=0.13.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace --user 0 $(protoImageName)

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh
.PHONY: proto-gen

debug: install
	colamosd start --pruning=nothing --evm.tracer=json --log_level trace \
                 --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545 \
                 --json-rpc.ws-address 0.0.0.0:8546

start: install
	colamosd start --pruning=nothing --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545  --json-rpc.ws-address 0.0.0.0:8546

docker: clean
	docker build --tag colamos-node -f ./Dockerfile .

reset: install init start

docker-test: install
	docker build --tag colamos-node -f ./Dockerfile.test .

clean:
	rm -rf $(BINS) $(CLEAN)

