package main

import (
	//"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"strconv"
)

var container *fyne.Container
var height int
var width int
var window fyne.Window
var a1 float64
var a2 float64
var a3 float64
var maxiter int
var noofitert *widget.Entry
var dc *gg.Context

func main() {
	application := app.New()
	window = application.NewWindow("Serpinski")
	container = fyne.NewContainer()
	renderbutton := widget.NewButton("Render", renderfunc)
	noofitert = widget.NewEntry()
	dc = gg.NewContext(1600, 900)
	window.SetPadded(false)
	container.Show()
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	window.Resize(fyne.NewSize(1600, 900))
	//container.Resize(fyne.NewSize(1600, 900))
	window.SetContent(widget.NewVBox(container, renderbutton, noofitert))
	window.ShowAndRun()
}

func drawtriangle(x, y, r, rotation float64, width float32, clr *color.RGBA) (fyne.Position, fyne.Position, fyne.Position) {
	x1 := x + r*math.Cos(rotation)
	y1 := y + r*math.Sin(rotation)
	x2 := x + r*math.Cos(rotation+gg.Radians(120))
	y2 := y + r*math.Sin(rotation+gg.Radians(120))
	x3 := x + r*math.Cos(rotation+gg.Radians(240))
	y3 := y + r*math.Sin(rotation+gg.Radians(240))
	line1 := canvas.NewLine(clr)
	line2 := canvas.NewLine(clr)
	line3 := canvas.NewLine(clr)
	line1.Position1.X = int(x1)
	line1.Position1.Y = int(y1)
	line1.Position2.X = int(x2)
	line1.Position2.Y = int(y2)
	line2.Position1.X = int(x2)
	line2.Position1.Y = int(y2)
	line2.Position2.X = int(x3)
	line2.Position2.Y = int(y3)
	line3.Position1.X = int(x3)
	line3.Position1.Y = int(y3)
	line3.Position2.X = int(x1)
	line3.Position2.Y = int(y1)
	drawline(x1, y1, x2, y2, clr.R, clr.G, clr.B, clr.A, width)
	drawline(x1, y1, x3, y3, clr.R, clr.G, clr.B, clr.A, width)
	drawline(x2, y2, x3, y3, clr.R, clr.G, clr.B, clr.A, width)
	container.AddObject(line1)
	container.AddObject(line2)
	container.AddObject(line3)
	return line1.Position1, line2.Position1, line3.Position1
}

func renderfunc() {
	clear()
	maxiter, _ = strconv.Atoi(noofitert.Text)
	height = dc.Height()
	width = dc.Width()
	triposy := float64(height / 2)
	triposx := float64(width / 2)
	length := 20 * float64(width) / 100
	recurse(triposx, triposy, -math.Pi/2, length, 1, maxiter)
}

func recurse(x, y, angle, length float64, width float32, iteration int) {
	clr := new(color.RGBA)
	clr.R = uint8(map1(float64(iteration), 1, float64(maxiter), 0, 255))
	clr.G = uint8(map1(angle, -math.Pi, math.Pi, 0, 255))
	clr.B = uint8(map1(y, 0, float64(dc.Height()), 0, 255))
	//fmt.Println(angle)
	position1, position2, position3 := drawtriangle(x, y, length, angle, width, clr)
	//	fmt.Println(gg.Degrees(angle))
	//	midpoint1x := float64(position1.X+position3.X) / 2
	//	midpoint1y := float64(position1.Y+position3.Y) / 2
	//	midpoint2x := float64(position2.X+position3.X) / 2
	//	midpoint2y := float64(position2.Y+position3.Y) / 2
	//	midpoint3x := float64(position1.X+position2.X) / 2
	//	midpoint3y := float64(position1.Y+position2.Y) / 2
	angle1 := math.Atan2(float64(position1.Y-position3.Y), float64(position1.X-position3.X))
	angle2 := math.Atan2(float64(position2.Y-position3.Y), float64(position2.X-position3.X))
	angle3 := math.Atan2(float64(position1.Y-position2.Y), float64(position1.X-position2.X))
	//	fmt.Println(gg.Degrees(angle1))
	//	fmt.Println(gg.Degrees(angle2))
	//	fmt.Println(gg.Degrees(angle3))
	fodderx1 := float64(position1.X+position3.X) / 2
	foddery1 := float64(position1.Y+position3.Y) / 2
	fodderx2 := float64(position2.X+position3.X) / 2
	foddery2 := float64(position2.Y+position3.Y) / 2
	fodderx3 := float64(position1.X+position2.X) / 2
	foddery3 := float64(position1.Y+position2.Y) / 2
	x1 := fodderx1 + (length/(3*1.35))*math.Cos(angle1-math.Pi/2)
	y1 := foddery1 + (length/(3*1.35))*math.Sin(angle1-math.Pi/2)
	x2 := fodderx2 + (length/(3*1.35))*math.Cos(angle2+math.Pi/2)
	y2 := foddery2 + (length/(3*1.35))*math.Sin(angle2+math.Pi/2)
	x3 := fodderx3 + (length/(3*1.35))*math.Cos(angle3+math.Pi/2)
	y3 := foddery3 + (length/(3*1.35))*math.Sin(angle3+math.Pi/2)
	if iteration > 0 {
		recurse(x1, y1, angle1-math.Pi/2+math.Pi*2/3, length/2, width/2, iteration-1)
		recurse(x2, y2, angle2+math.Pi/2+math.Pi*2/3, length/2, width/2, iteration-1)
		recurse(x3, y3, angle3+math.Pi/2+math.Pi*2/3, length/2, width/2, iteration-1)
	}
}

func clear() {
	for i := 0; i < len(container.Objects); i++ {
		container.Objects[i].Hide()
	}
}

func point(x, y float64, size float32) {
	clr := new(color.RGBA)
	clr.R = 50
	a := canvas.NewCircle(clr)
	a.Position1.X = int(x) - 5
	a.Position1.Y = int(y) - 5
	a.Position2.X = int(x) + 5
	a.Position2.Y = int(y) + 5
	a.Show()
	a.StrokeWidth = size
	container.AddObject(a)

}

func map1(value float64, istart float64, istop float64, ostart float64, ostop float64) float64 {
	return ostart + (ostop-ostart)*((value-istart)/(istop-istart))
}

var frameno int

func drawline(x1, y1, x2, y2 float64, r, g, b, a uint8, w float32) {
	//	fmt.Println(x1, y1, x2, y2)
	dc.DrawLine(x1, y1, x2, y2)
	dc.SetRGBA255(int(r), int(g), int(b), 150)
	dc.SetLineWidth(2)
	dc.Stroke()
	frameno++
	if frameno > 3113 {
		dc.SavePNG(strconv.Itoa(frameno) + ".png")
	}
}
