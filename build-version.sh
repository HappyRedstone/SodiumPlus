#!/bin/bash

# This script builds a specific version of the modpack.
#
# Usage:
# ./build-version.sh [minecraft version] [loader]
#
# Example:
# ./build-version.sh 1.21.10 Fabric

VERSION=$1
LOADER=$2

if [ -z "$VERSION" ] || [ -z "$LOADER" ]; then
    echo "Usage: ./build-version.sh [minecraft version] [loader]"
    exit 1
fi

go build -buildvcs=false -o bld ./builder
./bld version bundle $VERSION $LOADER
