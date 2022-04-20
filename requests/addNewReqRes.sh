#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

source "../../../../../scripts/bash_shared/convert_functions.sh"

[ $# -eq 0 ] && {
  echo "Usage: package (lowercase) Entity (CamelCase):"; exit 1;
}

namevar=$1
name=$2

fileExt=".go"
snakeName=$(convert_camel_to_snake_case_func "$name") # snake_case

fileName=$(to_lowercase "$snakeName")"$fileExt"

pack_name=$namevar

cat << EOF > $fileName
package $pack_name

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestFor$name(fabric string) *RequestFor$name {
  // @FIXME: Add correct API path
	return &RequestFor$name{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestFor$name struct {
	path string
}

func (req *RequestFor$name) Path() string {
	return req.path
}

func (req *RequestFor$name) Method() string {
	return http.MethodGet
}

func (req *RequestFor$name) Query() string {
	return ""
}

func (req *RequestFor$name) HasQueryParameter() bool {
	return false
}

func (req *RequestFor$name) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestFor$name) Payload() []byte {
	return nil
}

func (req *RequestFor$name) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseFor$name() *ResponseFor$name {
	return new(ResponseFor$name)
}

type ResponseFor$name struct {
  // @FIXME
	Field string `json:"field"`
}

func (r *ResponseFor$name) IsResponse() {}

func (r ResponseFor$name) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


EOF

## Generate build files & update bazel dependencies
command bazel run //:gazelle
## Build all sources
command bazel build //:build

git add .