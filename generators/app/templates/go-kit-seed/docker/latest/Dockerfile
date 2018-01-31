ARG ORG_NAME
ARG APP_NAME

FROM $ORG_NAME/$APP_NAME:install as builder
WORKDIR ./src
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /bin/app

FROM $ORG_NAME/$APP_NAME:base
COPY --from=builder /bin/app /
CMD [ "/app" ]
EXPOSE 443
