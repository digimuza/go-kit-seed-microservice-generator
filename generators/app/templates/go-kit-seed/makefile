include .env

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
	echo

reformat: 
	go fmt ./src

fetch-frontend:
	docker cp traefik:/cert .

generate-certificate: fetch-frontend
	openssl req -x509 -nodes -days 365 \
	-newkey rsa:2048 \
	-keyout ./cert/backend.key -out ./cert/backend.cert \
	-subj "/C=LT/ST=Lithuania/L=Kaunas/O=awpc/CN=be-keys-go""/C=LT/ST=Lithuania/L=Kaunas/O=$(ORG_NAME)/CN=$(APP_NAME)"

proto:
	./protoc pb/<%= appName %>.proto --go_out=plugins=grpc:./pkg

base: <% if(grpc) { %>proto <% } %>
	docker build \
		-f docker/base/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):base \
		docker/base

go: base
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/go/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):go \
		docker/go

install: go generate-certificate
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/install/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):install .

test-clean:
	docker-compose \
		-f docker/test/docker-compose.yml \
		rm -fsv || true

test-build: test-clean install
	docker-compose \
		-f docker/test/docker-compose.yml \
		build

test: test-build
	docker-compose \
		-f docker/test/docker-compose.yml \
		up 

# DEBUG


debug-build: install
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/debug/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):debug \
		docker/debug

debug: stop 
	docker rm -f $(APP_NAME)
	docker run \
		--security-opt=seccomp:unconfined \
		--network=awpc_webgateway \
		--env-file=.env \
		-p 2345:2345 --name $(APP_NAME) --rm $(ORG_NAME)/$(APP_NAME):debug

start-build: start-clean install
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/dev/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):dev \
		docker/dev

start-clean:
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		rm -fsv || true

start-create-db:
	docker exec \
		-itu postgres postgres \
		psql -c "SELECT create_db('$(DB_NAME)')"

init: start-build start-create-db debug-build

stop:
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		stop

start: 
	docker rm -f $(APP_NAME) || true
	docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		up

publish: install
	docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/prod/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):latest \
		docker/prod
	# docker push $(ORG_NAME)/$(APP_NAME):latest

deps:
	echo "  Dependencies: "
	echo
	echo "    * docker $(shell which docker > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * docker-compose $(shell which docker-compose > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * dep $(shell which dep > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo