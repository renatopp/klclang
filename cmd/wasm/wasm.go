//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/renatopp/klclang/internal"
)

func main() {
	runtime := internal.NewEvaluator()

	js.Global().Set("klc", js.FuncOf(func(this js.Value, args []js.Value) any {
		code := args[0].String()
		obj, err := internal.RunInRuntime([]byte(code), runtime)

		if err != nil {
			return js.ValueOf(err.Error())
		}
		return js.ValueOf(obj.String())
	}))

	select {} // keep running
}
