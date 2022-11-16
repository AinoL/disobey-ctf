#!/bin/sh

set -e

export GIN_MODE=release
/ctf > /var/log/ctf.stdout.log 2> /var/log/ctf.err.log &
