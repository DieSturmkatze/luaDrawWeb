package draw

import (
	"image"
	"image/color"
	"math"
)

type Vec2 struct {
	X, Y int
}

func Rect(x, y, w, h int, canvas *image.RGBA, color color.RGBA) {
	for xw := range w {
		canvas.Set(x+xw, y, color)
		canvas.Set(x+xw, y+h, color)
	}
	for yw := range h {
		canvas.Set(x, y+yw, color)
		canvas.Set(x+w, y+yw, color)
	}

}

func FillRect(pos Vec2, w, h int, canvas *image.RGBA, color color.RGBA) {
	for xw := range w {
		for yw := range h {
			canvas.Set(pos.X+xw, pos.Y+yw, color)
		}
	}
}

func LineCollection(points []Vec2, canvas *image.RGBA, color color.RGBA) {
	var prevPoint Vec2
	for i, point := range points {
		if i == 0 {
			prevPoint = point
			continue
		}

		Line(prevPoint, point, canvas, color)
		prevPoint = point
	}
}

func Line(p1, p2 Vec2, canvas *image.RGBA, color color.RGBA) {
	x1 := p1.X
	y1 := p1.Y
	x2 := p2.X
	y2 := p2.Y

	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			canvas.Set(x1, y, color)
		}
		return
	}

	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	m := dy / dx
	t := float64(y1) - m*float64(x1)

	for x := x1; x <= x2; x++ {
		y := roundInt(m*float64(x) + t)
		canvas.Set(x, y, color)
	}
}

func Circle(pos Vec2, rad int, canvas *image.RGBA, color color.RGBA) {
	x := pos.X
	y := pos.Y

	for i := range rad {
		px := i
		py := sqrt(rad*rad - i*i)

		canvas.Set(x+px, y+py, color)
		canvas.Set(x-px, y-py, color)
		canvas.Set(x-px, y+py, color)
		canvas.Set(x+px, y-py, color)

		py = i
		px = sqrt(rad*rad - i*i)

		canvas.Set(x+px, y+py, color)
		canvas.Set(x-px, y-py, color)
		canvas.Set(x-px, y+py, color)
		canvas.Set(x+px, y-py, color)
	}
}

func roundInt(num float64) int {
	return int(math.RoundToEven(num))
}

func sqrt(num int) int {
	return int(math.RoundToEven(math.Sqrt(float64(num))))
}

func abs(num int) int {
	return int(math.Abs(math.Sqrt(float64(num))))
}
