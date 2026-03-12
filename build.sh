#!/bin/bash

MAIN_GO_FILE="wasm-main.go"
OUT_DIR="_site"
WASM_LIB_DIR="Raylib-Go-Wasm"

set -e

if [ ! -e $MAIN_GO_FILE ]; then
    echo "error: $MAIN_GO_FILE is required"
    exit 1
fi

if [ ! -d $WASM_LIB_DIR ]; then
    echo "clone go wasm lib..."
    git clone https://github.com/BrownNPC/Raylib-Go-Wasm.git $WASM_LIB_DIR
    cd $WASM_LIB_DIR
    git checkout 27bc0271203c5039bbae9789b89d136c875b0976
    cd ..
fi

cp go.mod temp_go.mod
echo "replace github.com/gen2brain/raylib-go/raylib => ./$WASM_LIB_DIR/raylib" >> go.mod

go mod tidy

mkdir -p $OUT_DIR/rl

cp $WASM_LIB_DIR/index/rl/raylib.js   $OUT_DIR/rl/
cp $WASM_LIB_DIR/index/rl/raylib.wasm $OUT_DIR/rl/
cp $WASM_LIB_DIR/index/index.html     $OUT_DIR/
cp $WASM_LIB_DIR/index/index.js       $OUT_DIR/
cp $WASM_LIB_DIR/index/wasm_exec.js   $OUT_DIR/

echo "compile..."
GOOS=js GOARCH=wasm go build -o $OUT_DIR/main.wasm .

rm go.mod
mv temp_go.mod go.mod

echo "wasm successfully compiled!"
