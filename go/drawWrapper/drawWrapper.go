//go:build js && wasm

package drawWrapper

import (
	"fmt"
	"syscall/js"
)

const w, h = 640, 480

type Drawer struct {
	Interval int
}

func (d *Drawer) Println(text string) {
	fmt.Println("LUA ", text)
}

func (d *Drawer) Rect(x, y, w, h int, color string) {
	js.Global().Call("drawRect", x, y, w, h, color)
}

func (d *Drawer) FillRect(x, y, w, h int, color string) {
	js.Global().Call("drawFillRect", x, y, w, h, color)
}

func (d *Drawer) Line(x1, y1, x2, y2 int, color string) {
	js.Global().Call("drawLine", x1, y1, x2, y2, color)
}

func (d *Drawer) Circle(x, y, rad int, color string) {
	js.Global().Call("drawCircle", x, y, rad, color)
}

func (d *Drawer) PartialCircle(x, y, rad int, start, end float64, color string) {
	js.Global().Call("drawPartialCircle", x, y, rad, start, end, color)
}

func (d *Drawer) Text(x, y, size int, text, color string) {
	js.Global().Call("drawText", x, y, size, text, color)
}

func (d *Drawer) KeyDown(keycode int) bool {
	//fmt.Println(js.Global().Get("keysPressed").Get("D").String())
	return false
}

func (d *Drawer) ClearScreen() {
	js.Global().Call("clearCanvas")
}
