package drawWrapper

import (
	"fmt"
	"image"
	"image/color"

	"codeberg.org/mdanisch/luaDraw/draw"
	"github.com/zserge/fenster"
	"golang.org/x/image/colornames"
)

const w, h = 640, 480

type Drawer struct {
	Canvas          *image.RGBA
	Interval        int
	AutoClearScreen bool
	Win             fenster.Fenster
}

func (d *Drawer) Println(in string) {
	fmt.Println("[LUA] ", in)
}

func (d *Drawer) Rect(x, y, w, h int, color string) {
	draw.Rect(x, y, w, h, d.Canvas, stringToColor(color))
}

func (d *Drawer) FillRect(x, y, w, h int, color string) {
	draw.FillRect(draw.Vec2{X: x, Y: y}, w, h, d.Canvas, stringToColor(color))
}

func (d *Drawer) Line(x1, y1, x2, y2 int, color string) {
	draw.Line(draw.Vec2{X: x1, Y: y1}, draw.Vec2{X: x2, Y: y2}, d.Canvas, stringToColor(color))
}

func (d *Drawer) Circle(x, y, rad int, color string) {
	draw.Circle(draw.Vec2{X: x, Y: y}, rad, d.Canvas, stringToColor(color))
}

func (d *Drawer) ClearScreen() {
	d.Canvas = image.NewRGBA(image.Rect(0, 0, w, h))
}

func (d *Drawer) KeyDown(key int) bool {
	return d.Win.Key(byte(key))
}

func stringToColor(color string) color.RGBA {
	switch color {
	case "green":
		return colornames.Green
	case "red":
		return colornames.Red
	case "blue":
		return colornames.Blue
	case "black":
		return colornames.Black
	case "white":
		return colornames.White
	case "yellow":
		return colornames.Yellow
	case "purple":
		return colornames.Purple
	case "orange":
		return colornames.Orange
	default:
		return colornames.Black
	}
}
