ARG ORG_NAME
ARG APP_NAME

FROM $ORG_NAME/$APP_NAME:install

COPY entrypoint.sh /

RUN mv /cert ${APPPATH}/

RUN chmod +x /entrypoint.sh

CMD ["go", "test", "./..."]
