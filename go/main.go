package main

import (
	"image"
	imdraw "image/draw"
	"log"
	"time"

	"codeberg.org/mdanisch/luaDraw/cli"
	"codeberg.org/mdanisch/luaDraw/drawWrapper"

	lua "github.com/yuin/gopher-lua"
	"github.com/zserge/fenster"
	luar "layeh.com/gopher-luar"
)

var win fenster.Fenster

const w, h = 640, 480

func main() {
	luapath := cli.Setup()
	LS := lua.NewState()
	defer LS.Close()

	dw := &drawWrapper.Drawer{
		Canvas:          image.NewRGBA(image.Rect(0, 0, w, h)),
		Interval:        30,
		AutoClearScreen: true,
	}

	LS.SetGlobal("draw", luar.New(LS, dw))
	LS.DoFile(luapath)

	err := LS.CallByParam(lua.P{
		Fn:      LS.GetGlobal("init"),
		NRet:    1,
		Protect: true,
	})

	if err != nil {
		log.Fatalf("Your Lua init func has crashed with error: \n%s", err.Error())
	}

	//empty := image.NewRGBA(image.Rect(0, 0, w, h))
	win, err := fenster.New(w, h, "Thing")
	if err != nil {
		log.Fatal(err)
	}

	dw.Win = win
	defer win.Close()
	for win.Loop(time.Second / time.Duration(dw.Interval)) {

		if dw.AutoClearScreen {
			dw.Canvas = image.NewRGBA(image.Rect(0, 0, w, h))
		}

		err = LS.CallByParam(lua.P{
			Fn:      LS.GetGlobal("update"),
			NRet:    1,
			Protect: false,
		})
		if err != nil {
			log.Fatalf("Your Lua update func has crashed with error: \n%s", err.Error())
		}

		imdraw.Draw(win, win.Bounds(), dw.Canvas, image.Point{}, imdraw.Src)
	}
}
