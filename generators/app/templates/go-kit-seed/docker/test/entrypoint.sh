#!/bin/bash

dockerize \
    -wait tcp://${APP_NAME}.${DOMAIN}:443 \
    $@