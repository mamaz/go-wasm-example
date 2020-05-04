package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, playground")
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("printSimple", js.FuncOf(printSimple))
	fmt.Println("During channel waiting")
	<-done
}

func printSimple(this js.Value, args []js.Value) interface{} {
	return "Print sample"
}
