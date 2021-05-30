#! /bin/bash

CURRENT_VERSION=0.0.1
MODULE=github.com/joeyscat/oook
APP_NAME=oook

VERSION_SETTING=$MODULE'/internal/version'
GIT_COMMIT=$(git rev-parse --short HEAD || echo unsupported)

build_for() {
  go env -w GOARCH=$1
  go env -w GOOS=$2
  APP_NAME=$3
  OUTPUT_DIR='target/'$2'-'$1
  echo "$OUTPUT_DIR"

  if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir -p "$OUTPUT_DIR"
  fi

  go build -ldflags "-s -w \
    -X $VERSION_SETTING.Version=$CURRENT_VERSION \
    -X $VERSION_SETTING.GitCommit=$GIT_COMMIT \
    -X '$VERSION_SETTING.GoVersion=$(go version)' \
    -X '$VERSION_SETTING.BuildTime=$(date '+%Y-%m-%d %H:%M:%S')'" -o "$OUTPUT_DIR"'/'"$APP_NAME" ./main.go

  if [ $? != 0 ]; then
    echo "build failed"
    exit 1
  fi
}

# build_for 'amd64' 'windows' 'oook-.exe'
# build_for 'amd64' 'darwin' 'oook'
build_for 'amd64' 'linux' 'oook'

echo 'build success!'
echo 'version: ' $CURRENT_VERSION
echo 'git commit: ' "$GIT_COMMIT"
echo 'go version: ' "$(go version)"
