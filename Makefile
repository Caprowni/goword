.DEFAULT_TARGET=help
.PHONY: all
all: help

# VARIABLES
APP_NAME = goword
GO_VERSION ?= 1.14
GO_FILES = $(shell go list ./... | grep -v /vendor/)
PROJECT_PATH ?= github.com/lcaproni-pp/goword

VERSION ?= $(shell git describe --exact-match --tags 2>/dev/null)
COMMIT = $(shell git rev-parse HEAD | cut -c 1-6)
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

LDFLAGS = -ldflags "-s -w -X ${PROJECT_PATH}/cmd/version.Version=${VERSION} -X ${PROJECT_PATH}/cmd/version.Commit=${COMMIT} -X ${PROJECT_PATH}/cmd/version.BuildTime=${BUILD_TIME}"
RUN_CMD = docker run --rm -it -v "$(GOPATH):/go" -v "$(CURDIR)":/go/src/${PROJECT_PATH} -w /go/src/${PROJECT_PATH} golang:${GO_VERSION}-stretch

# COMMANDS
## install: install tool to your $GOPATH/bin directory
.PHONY: install
install:
	$(call blue, "# installing...")
	@go build $(LDFLAGS) -o $(GOPATH)/bin/goword

## help: Show this help message
.PHONY: help
help: Makefile
	@echo "${APP_NAME} - v${VERSION}"
	@echo " Choose a command run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^## //p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

# FUNCTIONS
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
