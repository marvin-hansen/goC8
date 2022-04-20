# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

BF=$(find . -name '*.bazel' | wc -l)
GF=$(find . -name '*.go' | wc -l)
PF=$(find . -name '*.proto' | wc -l)
gc=$(git rev-list --all --count master)

echo "Nr. Git commits: "$gc
echo "Nr. Go source files: "$GF
echo "Nr. Proto source files: "$PF
echo "Nr. Bazel Build files: "$BF

