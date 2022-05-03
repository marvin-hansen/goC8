# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# "---------------------------------------------------------"
# "-                                                       -"
# "-  Installs all dependencies required by make           -"
# "-                                                       -"
# "---------------------------------------------------------"

command bash --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install bash"
   exit
}
echo "* Bash installed"

command clang --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install clang"
   exit
}
echo "* Clang installed"

command echo "export CC=clang" >>~/.bashrc
echo "* Bash configured to use clang as CC for Bazel"

command curl --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install curl"
   exit
}
echo "* Curl installed"

command go version >/dev/null 2>&1 || {
   echo "Please install Golang"
       echo "https://go.dev/doc/install"
       exit
}
echo "* Golang installed"

echo ""
echo "==============================="
echo "All dependencies installed."
echo "==============================="
echo ""
