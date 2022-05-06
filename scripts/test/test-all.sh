# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# RUN test suites
echo "Testing Collection API"
command cd tests/collection
command go test -count=1 ./...

echo "Testing Document API"
command cd ../document
command go test -count=1 ./...

echo "Testing Graph API"
command cd ../graph
command go test -count=1 ./...

echo "Testing Index API"
command cd ../index
command go test -count=1 ./...

echo "Testing Key-Value API"
command cd ../kv
command go test -count=1 ./...

echo "Testing Query API"
command cd ../query
command go test -count=1 ./...

echo "Testing Query Worker API"
command cd ../query_worker
command go test -count=1 ./...
