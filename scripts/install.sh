#!/bin/bash

GAIN_SOURCE_PATH="/home/$USER/gain_source"
BIN_PATH="/home/$USER/gain/bin"
CURRENT_DIR="$pwd"

git clone https://github.com/andr-ll/gain $GAIN_SOURCE_PATH

mkdir -p $BIN_PATH

cd $GAIN_SOURCE_PATH

go build -o $BIN_PATH ./

rm -rf $GAIN_SOURCE_PATH

cd $CURRENT_DIR
