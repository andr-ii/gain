#!/bin/bash

punchy_SOURCE_PATH="/home/$USER/punchy_source"
BIN_PATH="/home/$USER/punchy/bin"
CURRENT_DIR="$pwd"

git clone https://github.com/andr-ii/punchy $punchy_SOURCE_PATH

mkdir -p $BIN_PATH

cd $punchy_SOURCE_PATH

go build -o $BIN_PATH ./

rm -rf $punchy_SOURCE_PATH

cd $CURRENT_DIR
