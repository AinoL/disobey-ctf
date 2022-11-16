#!/bin/sh

set -e

mkdir /images
/ctf > /var/log/ctf.stdout.log 2> /var/log/ctf.err.log &
