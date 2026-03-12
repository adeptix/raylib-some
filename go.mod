module raylib-some

go 1.25

//replace github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib

require github.com/gen2brain/raylib-go/raylib v0.55.1

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/sys v0.20.0 // indirect
)

//require github.com/BrownNPC/wasm-ffi-go v1.1.0 // indirect
require github.com/BrownNPC/wasm-ffi-go v1.2.0 // indirect
