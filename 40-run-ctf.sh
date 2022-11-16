#!/bin/sh

set -e

/ctf > /var/log/ctf.stdout.log 2> /var/log/ctf.err.log &
