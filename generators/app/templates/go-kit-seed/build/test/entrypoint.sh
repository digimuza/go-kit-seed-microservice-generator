#!/bin/bash

dockerize \
    -wait tcp://${APP_NAME}:443 \
    $@