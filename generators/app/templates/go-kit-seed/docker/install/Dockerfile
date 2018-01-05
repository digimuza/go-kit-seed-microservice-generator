ARG ORG_NAME
ARG APP_NAME
FROM $ORG_NAME/$APP_NAME:go

ENV APPPATH ${GOPATH}/src/app

RUN go get github.com/golang/dep \
 && cd $GOPATH/src/github.com/golang/dep \
 && go install ./...

WORKDIR ${APPPATH}

COPY /src ./src
COPY /Gopkg.lock ./Gopkg.lock
COPY /Gopkg.toml ./Gopkg.toml

RUN dep ensure

# WORKDIR ./src

# RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o ${GOPATH}/bin/app
