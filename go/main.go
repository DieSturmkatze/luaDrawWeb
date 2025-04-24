//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/DieSturmkatze/luaDrawWeb/go/drawWrapper"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func main() {
	fmt.Println("Loaded WASM")
	c := make(chan int)

	js.Global().Set("runLua", js.FuncOf(runLua))

	<-c
}

func runLua(this js.Value, inputs []js.Value) any {
	LS := lua.NewState()

	dw := &drawWrapper.Drawer{
		Interval:        30,
		AutoClearScreen: true,
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

	for js.Global().Get("stop").String() != "stop" {
		if dw.AutoClearScreen {
			dw.ClearScreen()
		}

		err = LS.CallByParam(lua.P{
			Fn:      LS.GetGlobal("update"),
			NRet:    1,
			Protect: true,
		})
		if err != nil {
			js.Global().Call("showError", err.Error)
			return ""
		}
	}
	return ""
}
