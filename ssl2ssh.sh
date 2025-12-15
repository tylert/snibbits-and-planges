#!/usr/bin/env bash

# Convert an OpenSSL private key to an OpenSSH private key.  Only supports
# ed25519 keys.  See https://security.stackexchange.com/a/268151 or
# https://security.stackexchange.com/questions/267711/how-can-i-convert-an-ed25519-key-in-pkcs8-to-openssh-private-key-format/268151#268151
# for details.

# Tools required:  bash, coreutils (base64, cat, dd, echo, printf, test), grep, openssl

# Invocation:
#   openssl genpkey -algorithm ed25519 | ./${0}
#   openssl genpkey -algorithm ed25519 -out key.txt ; ./${0} key.txt

set -euf

ssl_priv=$(cat ${1:+"${1}"})
pub64=$(echo "${ssl_priv}" | openssl pkey -pubout -outform der 2>/dev/null | dd bs=12 skip=1 status=none | base64)
test "${pub64}" || { echo 'Cannot get public key' >&2 ; exit 1 ; }
priv64=$(echo "${ssl_priv}" | grep -v '^-' | base64 -d | dd bs=16 skip=1 status=none | base64)

echo '-----BEGIN OPENSSH PRIVATE KEY-----'
{
    printf openssh-key-v1'\000\000\000\000\004'none'\000\000\000\004'none'\000\000\000\000\000\000\000\001\000\000\000'3
    printf '\000\000\000\013'ssh-ed25519'\000\000\000 '
    echo ${pub64} | base64 -d
    printf '\000\000\000'
    printf '\210\000\000\000\000\000\000\000\000'
    printf '\000\000\000\013'ssh-ed25519'\000\000\000 '
    echo ${pub64} | base64 -d
    printf '\000\000\000@'
    echo ${priv64} | base64 -d
    echo ${pub64} | base64 -d
    printf '\000\000\000\000\001\002\003\004\005'
} | base64
echo '-----END OPENSSH PRIVATE KEY-----'
