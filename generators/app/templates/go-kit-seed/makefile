# Variables

include .env

# Functions

define docker_build
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		--build-arg PROJECT_PATH=$(PROJECT_PATH) \
		-f build/${1}/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):${1} \
		${2}
endef

# Help

.SILENT:
help:
	echo
	echo "  \033[1;34m$(TITLE)\033[0m"
	echo "  ------------------------------------------------------------"
	echo
	echo "  Description: "
	echo
	echo "    Org: \033[1;34m$(ORG_NAME)\033[0m"
	echo "    App: \033[1;34m$(APP_NAME)\033[0m"
	echo
	echo "  Commands: "
	echo
	echo "    \033[1;32mhelp\033[0m    - show available commands"
	echo "    \033[1;32minit\033[0m    - setup project"
	echo "    \033[1;32mstart\033[0m   - start development"
	echo "    \033[1;32mtest\033[0m    - run tests"
	echo "    \033[1;32mdebug\033[0m   - run in debug mode"
	echo "    \033[1;32mpublish\033[0m - publish image"
	echo "    \033[1;32mdeps\033[0m    - check if all required dependencies are installed"
	echo "    \033[1;32mutils\033[0m   - run all utils commands"
	echo "    \033[1;32mclean\033[0m   - clean all docker images"
	echo
	echo "  Dependencies: "
	echo
	echo "    * docker $(shell which docker > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * docker-compose $(shell which docker-compose > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * glide $(shell which glide > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * protoc $(shell which protoc > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo
	echo

# Clean

clean: clean-container \
	clean-test \
	clean-start

clean-container:
	docker rm -f $(APP_NAME) || true

clean-test: 
	docker-compose \
		-f build/test/docker-compose.yml \
		rm -fsv || true

clean-start:
	docker-compose \
		-f build/dev/docker-compose.yml \
		-p $(APP_NAME) \
		rm -fsv || true
	docker-compose \
		-f build/dev/docker-compose.yml \
		-p $(APP_NAME) \
		build

# Utils

utils: utils-glide \
	utils-reformat \
	utils-proto

utils-glide:
	glide install

utils-reformat: 
	go fmt ./...

utils-proto:
	protoc pkg/pb/${PB_PATH}/${APP_NAME}.proto --go_out=plugins=grpc:./

# Build docker images

build: build-base \
	build-go \
	build-install \
	build-prod \
	build-dev \
	build-test

build-base:
	$(call docker_build,base,build/base)

build-go:
	$(call docker_build,go,build/go)

build-install:
	$(call docker_build,install,.)

build-prod:
	$(call docker_build,latest,build/latest)

build-dev:
	$(call docker_build,dev,build/dev)

build-test: 
	$(call docker_build,test,build/test)

build-debug: 
	$(call docker_build,debug,build/debug)

# Commands

init: utils build

stop:
	docker-compose \
		-f build/dev/docker-compose.yml \
		-p $(APP_NAME) \
		stop

start: clean-container
	docker-compose \
		-f build/dev/docker-compose.yml \
		-p $(APP_NAME) \
		up

publish: build-prod
	docker push $(ORG_NAME)/$(APP_NAME):latest

test: 
	go test ./pkg/... -cover

test-integration: clean-container clean-test build-install build-prod build-test
	docker-compose \
		-f build/test/docker-compose.yml \
		up 

debug: stop clean-container build-install build-debug
	docker run \
		--hostname=$(APP_NAME).$(DOMAIN) \
		--security-opt=seccomp:unconfined \
		--network=$(ORG_NAME)_webgateway \
		--env-file=.env \
      	-l traefik.protocol="https" \
      	-l traefik.backend="${APP_NAME}" \
      	-l traefik.frontend.rule="Host:${APP_NAME}.${DOMAIN}" \
      	-l traefik.priority="10" \
		-p 2345:2345 --name $(APP_NAME) --rm $(ORG_NAME)/$(APP_NAME):debug