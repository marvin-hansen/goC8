# Copyright (c) 2022. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# renames files in the working directory
function renameWorkflowFile() {
  name_var=$1
  name_lower=$2
  filename="workflow.go"

  search="NameDataManager"
  replace="$name_var""DataManager"
  rename_func "$filename" "$search"  "$replace"

  search="NameService"
  replace="$name_var""Service"
  rename_func "$filename" "$search"  "$replace"

  search="name"
  replace="$name_lower"
  rename_func "$filename" "$search"  "$replace"
}

function renameDataManager() {
  name_var=$1

  search="NameDataManager"
  replace="$name_var""DataManager"
  filename="deps.go"
  rename_func "$filename" "$search"  "$replace"

  filename="init.go"
  rename_func "$filename" "$search"  "$replace"

  echo "ok"
}

function renameServiceProvider() {
  name_var=$1

  search="NameService"
  replace="$name_var""Service"
  filename="deps.go"
  rename_func "$filename" "$search"  "$replace"

  filename="init.go"
  rename_func "$filename" "$search"  "$replace"

  echo "ok"
}

function renameImports() {
  name_lower=$1

  search="name"
  replace="$name_lower"
  filename="aNew.go"
  rename_func "$filename" "$search"  "$replace"

  filename="deps.go"
  rename_func "$filename" "$search"  "$replace"

  filename="init.go"
  rename_func "$filename" "$search"  "$replace"

  filename="init_verify.go"
  rename_func "$filename" "$search"  "$replace"

  filename="state.go"
  rename_func "$filename" "$search"  "$replace"

echo "ok"
}