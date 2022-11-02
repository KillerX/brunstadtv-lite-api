# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd sqlc gqlgen
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=btvplatform-api-lite
VERSION=0.0.1-alpha
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64 arm64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: build

all: clean sqlc gqlgen build_all

build:
	go build ${LDFLAGS} -o dist/${BINARY} ./src/cmd

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o dist/$(BINARY)-$(GOOS)-$(GOARCH) ./src/cmd )))

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

run:
	DB_PATH=$(DB_PATH) fresh

sqlc:
	sqlc generate

gqlgen:
	gqlgen

.PHONY: check clean build_all all run sqlc gqlgen
