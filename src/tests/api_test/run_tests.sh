#!/usr/bin/env bash
set -euo pipefail

go build -C ../../
mv ../../main ./kv

./kv &> /dev/null &
pid=$!

go test
ret=$?

kill "$pid"
rm ./kv
exit "$ret"
