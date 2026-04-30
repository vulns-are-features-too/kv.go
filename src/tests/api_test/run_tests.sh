#!/usr/bin/env bash
set -euo pipefail

go build -C ../../
mv ../../main ./kv

./kv &> /dev/null &
pid=$!

function cleanup() {
  rm ./kv
  kill "$pid"
}
trap cleanup EXIT

go test
ret=$?

exit "$ret"
