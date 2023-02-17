#!/usr/bin/env bash

go build \
    -a \
    -ldflags "-X main.Version=$(git describe --always --dirty --tags)" \
    -mod='vendor' \
    -trimpath='true'
