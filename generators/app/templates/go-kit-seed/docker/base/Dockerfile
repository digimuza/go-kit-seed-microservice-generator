FROM jwilder/dockerize

COPY entrypoint.sh /

RUN chmod +x /entrypoint.sh; \
    apk add --no-cache bash

ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 443 80