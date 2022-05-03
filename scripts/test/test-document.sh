# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# RUN KV Test suite
command cd tests/document
command go test -count=1 ./...

