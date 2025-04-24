//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/DieSturmkatze/luaDrawWeb/go/drawWrapper"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var LS *lua.LState
var dw *drawWrapper.Drawer
var done = make(chan struct{}) // Add this channel to block the main function

func main() {
	fmt.Println("Loaded WASM")

	js.Global().Set("runLua", js.FuncOf(runLua))
	js.Global().Set("runLuaInit", js.FuncOf(runLuaInit))
	js.Global().Set("goGetInterval", js.FuncOf(goGetInterval))
	<-done // Block the main function from exiting
}

func runLuaInit(this js.Value, inputs []js.Value) any {
	LS = lua.NewState()

	dw = &drawWrapper.Drawer{
		Interval: 1000,
	}

	LS.SetGlobal("draw", luar.New(LS, dw))

	luacode := inputs[0].String()
	fmt.Println(luacode)
	LS.DoString(luacode)

	err := LS.CallByParam(lua.P{
		Fn:      LS.GetGlobal("init"),
		NRet:    1,
		Protect: true,
	})

	if err != nil {
		js.Global().Call("showError", err.Error)
	}

	js.Global().Call("clearCanvas")
	return ""
}

func runLua(this js.Value, inputs []js.Value) any {
	//for js.Global().Get("stop").String() != "stop" {
	fmt.Println("Update")
	err := LS.CallByParam(lua.P{
		Fn:      LS.GetGlobal("update"),
		NRet:    1,
		Protect: true,
	})

	if err != nil {
		js.Global().Call("showError", err.Error)
		return ""
	}
	//}
	return ""
}

func goGetInterval(this js.Value, inputs []js.Value) any {
	return dw.Interval
}
