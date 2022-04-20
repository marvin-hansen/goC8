# Copyright (c) 2022. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Handy rename function.First parameter file, second what to search, third replace value.
# rename_func "$filename" "$search"  "$replace"
function rename_func() {
  filename=$1
  search=$2
  replace=$3
  if [[ $search != "" && $replace != "" ]]; then
    sed -i "" "s/$search/$replace/" "$filename"
fi
}

# Converts CameCase to snake_case
function convert_camel_to_snake_case_func() {
  echo "$1" | sed -r 's/([a-z0-9])([A-Z])/\1_\2/g'
}

# Converts UpperCase to lowercase
function to_lowercase() {
 echo "$1" | tr '[A-Z]' '[a-z]'
}

# Converts cap to CAP
function all_upper_case() {
 echo "$1" | tr '[a-z]' '[A-Z]'
}