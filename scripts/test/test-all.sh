# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# RUN test suites
command cd tests/collection
command go test -count=1 ./...

command cd ../document
command go test -count=1 ./...

command cd ../graph
command go test -count=1 ./...

command cd ../index
command go test -count=1 ./...

command cd ../kv
command go test -count=1 ./...

command cd ../query
command go test -count=1 ./...
