ARG ORG_NAME
ARG APP_NAME

FROM $ORG_NAME/$APP_NAME:install as builder
WORKDIR ./cmd/app
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /bin/server

FROM $ORG_NAME/$APP_NAME:base
COPY --from=builder /bin/server /
CMD [ "/server" ]
