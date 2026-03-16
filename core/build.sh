#!/bin/bash
set -e
cd "$(dirname "$0")"
echo "Building WASM..."
GOOS=js GOARCH=wasm go build -o wasm/highlighter.wasm highlighter.go
echo "Build complete: wasm/highlighter.wasm"