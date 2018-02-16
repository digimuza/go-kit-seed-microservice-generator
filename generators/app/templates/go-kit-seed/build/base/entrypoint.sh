#!/bin/bash

dockerize \
    -wait tcp://${DB_HOST}:${DB_PORT} \
    -wait tcp://${ZIPKIN_HOST}:${ZIPKIN_PORT} \
    $@