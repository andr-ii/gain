#!/bin/bash

PLT_SOURCE_PATH="/home/$USER/plt_source"
BIN_PATH="/home/$USER/plt/bin"
CURRENT_DIR="$pwd"

git clone https://github.com/andr-ll/plt $PLT_SOURCE_PATH

mkdir -p $BIN_PATH

cd $PLT_SOURCE_PATH

go build -o $BIN_PATH ./

rm -rf $PLT_SOURCE_PATH

cd $CURRENT_DIR
