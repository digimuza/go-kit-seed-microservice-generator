[![N|Solid](https://www.passcamp.com/wp-content/uploads/2017/11/PassCampColor.png)](https://pass.camp/start/login)

# PassCamp <%= serviceName %> microservice
Microservice that store buckets.

# Development best practices - (GO KIT)
[Go kit](https://github.com/go-kit/kit) is a programming toolkit for building microservices (or elegant monoliths) in Go.

## Dependencies

  - [Docker](https://docs.docker.com/) - [INSTALATION GUIDE](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04)
  - [docker-compose](https://docs.docker.com/compose/) - 1.18.0
  - [glide](https://github.com/Masterminds/glide) - go get -u github.com/Masterminds/glide
  - [google/protobuf](https://github.com/google/protobuf/releases) - put bin && includes files to /usr/local
      DOWNLOAD - realese package for your mashine - [3.5.1-linux-x86_64.zip](https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip)
      EXTRACT zip package and move bin && includes folders to /usr/local

  - [golang/protobuf](https://github.com/golang/protobuf) - go get -u github.com/golang/protobuf/protoc-gen-go



### Install go 

[How to install Go language](https://golang.org/doc/install)

```sh
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

### Prepare for setup

  - Create folders $GOPATH/src/pas/dev.adeoweb.biz/pas
  - Clone git repository git clone https://dbarzdys@dev.adeoweb.biz:8453/scm/pas/<%= appName %>.git

### Check dependencies
```sh
$ make help
```

### Initial Setup 
```sh
$ make init
```

### How to install new dependency?
It's easy just use:
```sh
$ glide update
```

### Start 
```sh
$ make start
```
This command will spin up development environment.
- localhost:3000 - letmegrpc
- localhost:5050 - adminer
- localhost:9090 - prometheus
- localhost:9411 - zipkin
- localhost:5000 - grafana

### Run all test in an isolated environment. 
```sh
$ make tests
```

### Debug 
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

