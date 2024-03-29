#!/usr/bin/env bash

# tool='mybinary'
variable='Version'

version="${1}"
if [ ! -z "${version}" ]; then
    version="$(git describe --always --dirty --tags)"
fi

# GOOS=${os} GOARCH=${arch}
CGO_ENABLED=0 go build \
    -a \
    -ldflags "-X main.${variable}=${version} -s -w" \
    -mod='vendor' \
    -trimpath
    # -o ${tool}
# zip ${tool}-${os}-${arch}-${version}.zip ${tool}

# syft packages ${tool} -o spdx-json=sbom.spdx.json
