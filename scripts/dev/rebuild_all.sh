#
# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com
#

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

bazel run //:gazelle

# Convert mod dependencies into bazel dependencies
bazel run //:gazelle -- update-repos -from_file=go.mod

# Update all build files & dependencies
bazel run //:gazelle

# Build all sources
bazel build //:build
