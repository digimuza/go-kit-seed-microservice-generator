#!/bin/bash

dockerize \
    -wait tcp://${DB_HOST}:${DB_PORT} \
    $@