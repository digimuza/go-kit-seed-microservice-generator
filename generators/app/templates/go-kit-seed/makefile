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

proto:
	./protoc pb/<%= appName %>.proto --go_out=./pkg

base: <% if(grpc) { %>proto <% } %>
	sudo docker build \
		-f docker/base/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):base \
		docker/base

go: base
	sudo docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/go/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):go \
		docker/go

install: go
	sudo docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/install/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):install .

test-clean:
	sudo docker-compose \
		-f docker/test/docker-compose.yml \
		rm -fsv || true

test-build: test-clean install
	sudo docker-compose \
		-f docker/test/docker-compose.yml \
		build

test: test-build
	sudo docker-compose \
		-f docker/test/docker-compose.yml \
		up 

start-build: start-clean install
	sudo docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/dev/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):dev \
		docker/dev

start-clean:
	sudo docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		rm -fsv || true

start-create-db:
	sudo docker exec \
		-itu postgres postgres \
		psql -c "SELECT create_db('$(DB_NAME)')"

start: start-build start-create-db
	sudo docker-compose \
		-f docker/dev/docker-compose.yml \
		-p $(APP_NAME) \
		up

publish: install
	sudo docker build \
		--build-arg ORG_NAME=$(ORG_NAME) \
		--build-arg APP_NAME=$(APP_NAME) \
		-f docker/prod/Dockerfile \
		-t $(ORG_NAME)/$(APP_NAME):latest \
		docker/prod
	sudo docker push $(ORG_NAME)/$(APP_NAME):latest

deps:
	echo "  Dependencies: "
	echo
	echo "    * docker $(shell which docker > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * docker-compose $(shell which docker-compose > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo "    * dep $(shell which dep > /dev/null || echo '- \033[31mNOT INSTALLED\033[37m')"
	echo