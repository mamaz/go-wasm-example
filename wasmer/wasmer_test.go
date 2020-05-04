package wasmer

import (
	"fmt"
	"testing"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func TestWasmerRunWasmBinary(t *testing.T) {
	bytes, _ := wasmer.ReadBytes("../wasm/main.wasm")
	instance, _ := wasmer.NewInstance(bytes)
	defer instance.Close()

	printSample := instance.Exports["printSample"]
	result, _ := printSample()
	fmt.Println(result) // DOESN't WORK  YET
}
