package main

import "syscall/js"

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	ch := make(chan struct{})

	goFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		n := args[0].Int()
		return fib(n)
	})
	defer goFunc.Release()
	js.Global().Set("goFunc", goFunc)

	<-ch
}
