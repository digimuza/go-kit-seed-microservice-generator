ARG ORG_NAME
ARG APP_NAME
FROM $ORG_NAME/$APP_NAME:install

ENV DELVE_VERSION=0.12.2

RUN apk --no-cache add tini \
    && apk --no-cache add --virtual build-dependencies git \
    && go get github.com/derekparker/delve/cmd/dlv \
    && cd $GOPATH/src/github.com/derekparker/delve \
    && git checkout v${DELVE_VERSION} \
    && go install github.com/derekparker/delve/cmd/dlv

COPY debug.sh /

RUN chmod +x /debug.sh
    
CMD ["/debug.sh"]

EXPOSE 2345 443