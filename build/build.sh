#!/usr/bin/env bash

# script directory
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# cd to ./src
cd $SCRIPTDIR/../src

# build the binary
go build -o ../bin/adventofcode2020

# copy input assets
mkdir ../bin/inputs >/dev/null 2>&1
yes | cp ./inputs/* ../bin/inputs -rf