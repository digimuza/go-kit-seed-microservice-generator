[![N|Solid](https://www.passcamp.com/wp-content/uploads/2017/11/PassCampColor.png)](https://pass.camp/start/login)

# PassCamp <%= appName %> microservice
Microservice that store users private and public keys.

# Development best practices - (GO KIT)
[Go kit](https://github.com/go-kit/kit) is a programming toolkit for building microservices (or elegant monoliths) in Go.

## Dependencies

  - [Docker](https://docs.docker.com/) - [INSTALATION GUIDE](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04)
  - [docker-compose](https://docs.docker.com/compose/) - 1.18.0
  - [golang/dep](https://github.com/golang/dep) - go get -u github.com/golang/dep/cmd/dep
  - [google/protobuf](https://github.com/google/protobuf/releases) - Put bin folder to /usr/local
  - [golang/protobuf](https://github.com/golang/protobuf) - go get -u github.com/golang/protobuf/protoc-gen-go
  - [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master) - 1.0 

### Install go 

[How to install Go language](https://golang.org/doc/install)

```sh
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

### Setup /etc/hosts file
```sh
127.0.0.1       *******
127.0.1.1       *******
127.0.0.1       *******
127.0.0.1       be-buckets.passcamp.doc
127.0.0.1       postgres
127.0.0.1       *******
127.0.1.1       *******
127.0.0.1       *******
```

### Prepare for setup

  - Create folder <awpc> in $GOPATH/src/ 
  - Install and start [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master)

### Initial Setup 
```sh
$ make init
```

### Check dependencies
```sh
$ make help
```

### Build microservice 
```sh
$ make build
```

### How to install new dependency?
It's easy just use:
```sh
$ go get github.com/asaskevich/govalidator
```

### Start 
Make sure that [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master) is running:
```sh
$ make start
```

### Run single test || all tests from vscode editor
Make sure that [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master) and be-keys microservice is running. Use vscode run test command.

### Run all test in an isolated environment. 
Make sure that [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master) is running:
```sh
$ make tests
```

### Debug 
Make sure that [awpc-dev-stack](https://dev.adeoweb.biz:8453/projects/PAS/repos/dev-stack/browse?at=refs/heads/master) is running: After that use vscode debug with Remote configuration
```sh
$ make debug
```

### Publish and deploy code.
Make sure that your account has correct permissions
```sh
$ make publish
```

## Useful links
- [GO-KIT-MICROSERVICE GENERATOR](https://github.com/digimuza/go-kit-seed-microservice-generator) - Generate initial microservice boilerplate
- [Gorm](http://jinzhu.me/gorm/) - The fantastic ORM library for Golang, aims to be developer friendly.
- [govalidator](https://github.com/asaskevich/govalidator) - A package of validators and sanitizers for strings, structs, and collections.
- [Proto-buffers](https://developers.google.com/protocol-buffers/docs/overview) - New communication protocol

