# Variables

include .env

# Functions

define docker_build
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/${1}/Dockerfile \
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
	echo "    \033[1;32mstart\033[0m   - start development"
	echo "    \033[1;32mtest\033[0m    - run tests"
	echo "    \033[1;32mpublish\033[0m - publish image"
	echo "    \033[1;32mdeps\033[0m    - check if all required dependencies are installed"
	echo
	echo "  Dependencies: "
	echo
	echo "    * docker $(shell which docker > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * docker-compose $(shell which docker-compose > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * dep $(shell which dep > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
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
		-f docker/test/docker-compose.yml \
		rm -fsv || true

clean-start:
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		rm -fsv || true

# Utils

utils: utils-dep \
	utils-reformat \
	utils-copy-cert \
	<% if(grpc){ %>utils-proto \<% } %> 
	utils-create-db

utils-dep:
	dep ensure

utils-reformat: 
	go fmt ./src

utils-copy-cert:
	docker cp traefik:/cert .

utils-gen-cert:
	openssl req -x509 -nodes -days 365 \
	-newkey rsa:2048 \
	-keyout ./cert/backend.key -out ./cert/backend.cert \
	-subj "/C=LT/ST=Lithuania/L=Kaunas/O=awpc/CN=be-keys-go""/C=LT/ST=Lithuania/L=Kaunas/O=$(ORG_NAME)/CN=*"

utils-proto:
	./protoc pb/$(APP_NAME).proto --go_out=plugins=grpc:./pkg

utils-create-db:
	docker exec \
		-itu postgres postgres \
		psql -c "SELECT create_db('$(DB_NAME)')"

# Build docker images

build: build-base \
	build-go \
	build-install \
	build-prod \
	build-dev \
	build-test

build-base:
	$(call docker_build,base,docker/base)

build-go:
	$(call docker_build,go,docker/go)

build-install:
	$(call docker_build,install,.)

build-prod:
	$(call docker_build,latest,docker/latest)

build-dev:
	$(call docker_build,dev,docker/dev)

build-test: 
	$(call docker_build,test,docker/test)

build-debug: 
	$(call docker_build,debug,docker/debug)

# Commands

init: utils build

stop:
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		stop

start: utils-reformat clean-container
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		up

publish: build-prod
	docker push $(ORG_NAME)/$(APP_NAME):latest

test: clean-container clean-test build-install build-prod build-test
	docker-compose \
		-f docker/test/docker-compose.yml \
		up 

debug: stop clean-container build-install build-debug
	docker run \
		--security-opt=seccomp:unconfined \
		--network=awpc_webgateway \
		--env-file=.env \
      	-l traefik.protocol="https" \
      	-l traefik.backend="${APP_NAME}" \
      	-l traefik.frontend.rule="Host:${APP_NAME}.${DOMAIN}" \
      	-l traefik.priority="10" \
		-p 2345:2345 --name $(APP_NAME) --rm $(ORG_NAME)/$(APP_NAME):debug
