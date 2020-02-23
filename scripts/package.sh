#!/bin/bash
set -e

BASE_PACKAGE=build/packages
rm -Rf $BASE_PACKAGE
mkdir -p $BASE_PACKAGE

build() {
  # $1 -> operating system
  # $2 -> architecture
  # $3 -> OS alias, used in the output file name
  # $4 -> Optional extension with ".", e.g.: .exe
  PACKAGE_FILE=$BASE_PACKAGE/spa-server_$3_$2$4
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -a -installsuffix cgo -o $PACKAGE_FILE ./cmd
}

build darwin amd64 mac
build linux amd64 linux
build windows amd64 win .exe
