package main

import (
	"syscall/js"
)

func sumInt(array js.Value, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += array.Index(i).Int()
	}
	return sum
}

func main() {
	ch := make(chan struct{})

	goFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		n := args[1].Int()
		return sumInt(args[0], n)
	})
	defer goFunc.Release()
	js.Global().Set("goFunc", goFunc)

	<-ch
}
