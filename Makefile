APP=neosync
ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
COMMIT=$(shell git rev-parse HEAD | cut -c 1-8)
LDFLAGS="-w -s -X $(BUILD_INFO_PKG).BuildTime=$$(TZ=UTC date '+%FT%T') -X neosync.AppVersion=$$(git rev-parse HEAD | cut -c 1-8) -X $(BUILD_INFO_PKG).VCSRef=$$(git rev-parse --abbrev-ref HEAD)"

BIN_DIR=build

.PHONY: build dev

build:
	go build -o $(BIN_DIR)/$(APP) -ldflags $(LDFLAGS) ./cmd/.

build-docker:
	docker build -f ./deploy/Dockerfile . -t $(APP):$$(git rev-parse HEAD | cut -c 1-8)

dev:
	which CompileDaemon || (GO111MODULE=off go get github.com/githubnemo/CompileDaemon)
	HOST_IP=127.0.0.1 CompileDaemon -color=true -log-prefix=false --build="make build" -exclude-dir="build" --command="./$(BIN_DIR)/$(APP)"

dev-up:
	docker compose -f ./deploy/docker-compose.yml -p neosync up -d


dev-down:
	docker compose -f ./deploy/docker-compose.yml -p neosync down --volumes