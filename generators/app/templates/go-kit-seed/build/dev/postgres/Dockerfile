FROM postgres:9-alpine

ENV POSTGRES_USERNAME postgres
ENV POSTGRES_PASSWORD postgres
ENV PGDATA /var/lib/postgresql/data/pgdata

ADD *.sql /docker-entrypoint-initdb.d/
