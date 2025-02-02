#!/bin/bash
set -e

if [ -z "$1" ]; then
    echo "Usage: ./build.sh <version>"
    exho "Example: ./build.sh 1.0.0"
    exit 1
fi
VERSION=$1

PACKAGE="github.com/Dias1c/aws-ses-bulk-emails"
PROGRAM_NAME="aws_ses_bulk_emails"
PLATFORMS=("windows/amd64" "linux/amd64" "darwin/amd64" "darwin/arm64")

mkdir -p dist

for PLATFORM in "${PLATFORMS[@]}"; do
    OS=${PLATFORM%/*}
    ARCH=${PLATFORM#*/}

    OUTPUT="$PROGRAM_NAME-v$VERSION-$OS-$ARCH"
    if [ "$OS" == "windows" ]; then
        OUTPUT+=".exe"
    fi

    echo "Building for $OS/$ARCH..."
    GOOS=$OS GOARCH=$ARCH go build  -ldflags "-X '$PACKAGE/internal/configs.Vesrion=$VERSION'" -o dist/$OUTPUT ./cmd/quick
done

echo "Build completed!"

