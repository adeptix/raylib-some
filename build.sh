#!/bin/bash

MAIN_GO_FILE="wasm-main.go"
OUT_DIR="_site"
WASM_LIB_DIR="Raylib-Go-Wasm"

set -e

if [ ! -e $MAIN_GO_FILE ]; then
    echo "error: $WASM_LIB_DIR is required"
    exit 1
fi

if [ ! -d $WASM_LIB_DIR ]; then
    echo "clone go wasm lib..."
    git clone https://github.com/BrownNPC/Raylib-Go-Wasm.git
fi

cp go.mod temp_go.mod
echo "replace github.com/gen2brain/raylib-go/raylib => ./$WASM_LIB_DIR/raylib" >> go.mod
echo "require github.com/BrownNPC/wasm-ffi-go v1.2.0 // indirect" >> go.mod

go mod download

mkdir -p $OUT_DIR/rl

cp $WASM_LIB_DIR/index/rl/raylib.js   $OUT_DIR/rl/
cp $WASM_LIB_DIR/index/rl/raylib.wasm $OUT_DIR/rl/
cp $WASM_LIB_DIR/index/index.html     $OUT_DIR/
cp $WASM_LIB_DIR/index/index.js       $OUT_DIR/
cp $WASM_LIB_DIR/index/wasm_exec.js   $OUT_DIR/

echo "compile..."
GOOS=js GOARCH=wasm go build -o $OUT_DIR/main.wasm wasm-main.go

rm go.mod
mv temp_go.mod go.mod

echo "wasm successfully compiled!"
