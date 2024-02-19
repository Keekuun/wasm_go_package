//go:build js && wasm
// +build js,wasm

package main

import (
	_ "embed"
	"errors"
	"syscall/js"
	"wasm_go_package/utils"
)

// 注入说明文档
//
//go:embed info.txt
var info []byte

func readme(this js.Value, args []js.Value) any {
	return string(info)
}

func add(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return utils.PackResult(false, errors.New("参数不能为空！"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackResult(true, nil, x+y)
}

func sub(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return utils.PackResult(false, errors.New("参数不能为空！"), nil)
	}
	x := args[0].Float()
	y := args[1].Float()
	return utils.PackResult(true, nil, x-y)
}

func main() {
	js.Global().Set("readme", js.FuncOf(readme))
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("sub", js.FuncOf(sub))
	// 阻塞挂起
	<-make(chan any)

	//fmt.Println(readme())
	//fmt.Println(add([]float64{1, 2}))
	//fmt.Println(sub([]float64{1, 2}))
}
