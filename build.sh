#!/bin/bash

platforms=("windows/amd64" "darwin/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="builds/DiscordBuilder-${GOOS}-x86_64"
    if [ "$GOOS" = "windows" ]; then
        output_name+=".exe"
    fi
    echo "Building for $GOOS/$GOARCH..."
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$output_name" main.go
    echo "Build complete for $GOOS/$GOARCH"
    printf "\n"
done
