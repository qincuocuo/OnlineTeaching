#!/bin/sh

# shellcheck disable=SC2112
function main() {
#  git pull origin master
  go mod vendor
  docker build -t sun:1.0 -f build/dockerfile/Dockerfile .
  docker save sun:1.0 -o build/images/sun.tar
  scp build/images/sun.tar guest@121.199.167.227:/home/guest/onlineTeaching/images
}

main
