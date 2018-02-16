ARG ORG_NAME
ARG APP_NAME

FROM $ORG_NAME/$APP_NAME:install

RUN go get github.com/loomnetwork/fresh

COPY runner.conf .

CMD ["fresh"]
