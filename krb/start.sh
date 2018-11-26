#!/usr/bin/env bash
set -xe

echo 127.0.0.1 my.ex >> /etc/hosts
krb5kdc

exec /usr/local/bin/docker-entrypoint.sh "$@"
