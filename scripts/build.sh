#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

#suggest go version 1.20.2

if ! [[ "$0" =~ scripts/build.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
BUILD_PATH=$ROOT_PATH/build/app
SRC_PATH=$ROOT_PATH/src/main/main.go
MODULE=GoWeb

if [ "$ROOT_PATH" = "$PWD" ]; then
  echo "running from repsitory root"
else
  echo "root path and pwd path is not equal"
  exit 255
fi

rm -rf build

git_commit=$(git rev-parse HEAD)
echo $git_commit

# Download dependencies
echo "Downloading dependencies..."
go mod tidy
go mod download

echo "start build app"
go build -ldflags "-X $MODULE/src/version.VerionGit=$git_commit" -o "$BUILD_PATH" "$SRC_PATH"
echo "build success"

cp -r conf/ build/
# mkdir build/logs
