package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Get("console").Call("log", "WASM loaded")
	fmt.Println("Gaslands.tools WASM module initialized")

	// Block forever — keep WASM alive for JS calls.
	select {}
}
