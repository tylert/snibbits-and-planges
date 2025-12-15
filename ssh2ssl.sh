#!/usr/bin/env bash

# Convert an OpenSSH private key to an OpenSSL one.  See
# https://gist.github.com/csobankesmarki/ea5e46a5af68623f1d9b442af1528225?permalink_comment_id=4421070#gistcomment-4421070
# for details.

# Tools required:  bash, cat, coreutils (base64, dd, printf, tr), grep, openssl

# Invocation:
#   ssh-keygen -f ssh.key -t ed25519 -q -C '' -N '' ; cat ssh.key | ./${0} > ssl.key

cat | (
  printf \\x30\\x2e\\x02\\x01\\x00\\x30\\x05\\x06\\x03\\x2b\\x65\\x70\\x04\\x22\\x04\\x20
  grep -Ev '^-' | tr -d '\n' | base64 -d | dd bs=1 skip=161 count=32 status=none
) | openssl pkey
