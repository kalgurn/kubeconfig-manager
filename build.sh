#!/bin/bash

PACKAGE_NAME="kubeconfig-manager"
MAIN_PACKAGE="cmd/${PACKAGE_NAME}/main.go"
COMMAND_PACKAGE="github.com/kalgurn/${PACKAGE_NAME}/internal/cmd"

BINARY_NAME="${PACKAGE_NAME}-${GOOS}-${GOARCH}"

GITHASH=$(git rev-parse --short HEAD)
DATE=$(date -u)

VERSION=$1

go build -o ${BINARY_NAME} -ldflags "-X '${COMMAND_PACKAGE}.Version=${VERSION}' -X '$COMMAND_PACKAGE.BuildDate=${DATE}' -X '${COMMAND_PACKAGE}.CommitSHA=${GITHASH}'" ${MAIN_PACKAGE}

zip ${BINARY_NAME}.zip ${BINARY_NAME} > /dev/null