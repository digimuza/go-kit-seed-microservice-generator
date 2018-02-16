ARG ORG_NAME
ARG APP_NAME
FROM $ORG_NAME/$APP_NAME:go

ARG ORG_NAME
ARG APP_NAME
ARG PROJECT_PATH

ENV APPPATH ${GOPATH}/src/${PROJECT_PATH}/${APP_NAME}

RUN go get github.com/golang/dep \
 && cd $GOPATH/src/github.com/golang/dep \
 && go install ./...

WORKDIR ${APPPATH}

COPY /cert /cert

COPY /cmd ./cmd
RUN ls cmd/app
COPY /pkg ./pkg
COPY /internal ./internal
COPY /vendor ./vendor

COPY /glide.lock ./glide.lock
COPY /glide.yaml ./glide.yaml