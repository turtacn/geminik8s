#!/bin/bash

# build.sh: A script to build the geminik8s binary for production.

set -e # Exit immediately if a command exits with a non-zero status.
set -o pipefail # Return value of a pipeline is the value of the last command to exit with a non-zero status

# Set the project root directory
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT_DIR"

# Define binary name and main package path
BINARY_NAME="gemin_k8s"
MAIN_PATH="github.com/turtacn/geminik8s/cmd/gemin_k8s"
OUTPUT_DIR="./dist/linux_amd64"

# Get version information from git
VERSION=$(git describe --tags --always --dirty)
COMMIT=$(git rev-parse HEAD)
DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
BUILT_BY="build.sh"

# Define the path for the version variables in the Go code
VERSION_PKG="github.com/turtacn/geminik8s/internal/app/cli"

# Setup Go LDFLAGS to inject version info
LDFLAGS=(
  "-X '${VERSION_PKG}.Version=${VERSION}'"
  "-X '${VERSION_PKG}.Commit=${COMMIT}'"
  "-X '${VERSION_PKG}.Date=${DATE}'"
  "-X '${VERSION_PKG}.BuiltBy=${BUILT_BY}'"
)

# Join the array elements with spaces
LDFLAGS_STR="${LDFLAGS[*]}"

echo "================================================================================"
echo " Building geminik8s"
echo "================================================================================"
echo " > Version:    ${VERSION}"
echo " > Commit:     ${COMMIT}"
echo " > Build Date: ${DATE}"
echo " > Main Path:  ${MAIN_PATH}"
echo " > Output Dir: ${OUTPUT_DIR}"
echo "================================================================================"

# Build the binary for Linux
mkdir -p "$OUTPUT_DIR"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS_STR}" -o "${OUTPUT_DIR}/${BINARY_NAME}" "${ROOT_DIR}/cmd/gemin_k8s/main.go"

echo "Build complete. Binary is at ${OUTPUT_DIR}/${BINARY_NAME}"

#Personal.AI order the ending
