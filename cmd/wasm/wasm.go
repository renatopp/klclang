//go:build js && wasm

package main

import "syscall/js"

func CallMe(this js.Value, args []js.Value) any {
	println("CallMe called!")
	return nil
}

func main() {
	println("Hello, from Go!")
	js.Global().Set("CallMe", js.FuncOf(CallMe))
	select {} // keep running
}
