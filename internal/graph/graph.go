package graph

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"path/filepath"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/smackem/goplot/internal/calc"
)

func init() {
	draw2d.SetFontFolder(filepath.FromSlash("resource/font"))
}

func DrawPng() image.Image {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	return dest
}

func PlotPng(xs []calc.Number, ys []calc.Number, minY, maxY float64) image.Image {
	if len(xs) <= 0 || len(xs) != len(ys) {
		panic("xs and ys must be of equal length, both greater than zero")
	}

	bounds := image.Rect(0, 0, 800, 600)
	if minY == 0 && maxY == 0 {
		minY, maxY = minMax(ys)
	}
	minX, maxX := float64(xs[0]), float64(xs[len(xs)-1])
	ratioX := float64(bounds.Dx()) / float64(maxX-minX)
	ratioY := float64(bounds.Dy()) / float64(maxY-minY)

	log.Printf("PlotPng: x=%g:%g y=%g:%g ratio=%g,%g\n", minX, maxX, minY, maxY, ratioX, ratioY)

	dest := image.NewRGBA(bounds)
	gc := draw2dimg.NewGraphicContext(dest)

	// fill with white
	draw.Draw(dest, bounds, image.White, image.ZP, draw.Over)

	gc.SetStrokeColor(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
	gc.SetLineWidth(1)

	// transform to center 0,0
	gc.Translate(math.Abs(minX*ratioX)+0.5, math.Abs(maxY*ratioY)+0.5)

	// draw x and y axis
	gc.MoveTo(minX*ratioX, 0)
	gc.LineTo(maxX*ratioX, 0)
	gc.Stroke()

	gc.MoveTo(0, -minY*ratioY)
	gc.LineTo(0, -maxY*ratioY)
	gc.Stroke()

	// draw axis labels
	gc.SetFontData(draw2d.FontData{Name: "luxi", Family: draw2d.FontFamilySans, Style: draw2d.FontStyleBold})
	gc.SetFontSize(10)
	gc.SetFillColor(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})

	maxXStr := fmt.Sprintf("%.2f", maxX)
	left, top, right, bottom := gc.GetStringBounds(maxXStr)
	gc.FillStringAt(maxXStr, maxX*ratioX-(right-left)-1, bottom-top+1)

	minXStr := fmt.Sprintf("%.2f", minX)
	left, top, right, bottom = gc.GetStringBounds(minXStr)
	gc.FillStringAt(minXStr, minX*ratioX+1, bottom-top+1)

	maxYStr := fmt.Sprintf("%.2f", maxY)
	left, top, right, bottom = gc.GetStringBounds(maxYStr)
	gc.FillStringAt(maxYStr, 1, -maxY*ratioY+(bottom-top)+1)

	minYStr := fmt.Sprintf("%.2f", minY)
	left, top, right, bottom = gc.GetStringBounds(minYStr)
	gc.FillStringAt(minYStr, 1, -minY*ratioY-2)

	// draw curve
	gc.SetLineWidth(3)
	gc.SetStrokeColor(color.RGBA{0, 0, 0, 0xff})
	gc.MoveTo(float64(xs[0])*ratioX, -float64(ys[0])*ratioY)
	for i := 1; i < len(xs); i++ {
		x, y := float64(xs[i])*ratioX, -float64(ys[i])*ratioY
		gc.LineTo(x, y)
	}
	gc.Stroke()

	return dest
}

func fillRect(dest draw.Image, x, y, width, height float64, rgb color.Color) {
	rect := image.Rect(int(x+0.5), int(y+0.5), int(x+width+0.5), int(y+height+0.5))
	draw.Draw(dest, rect, image.NewUniform(rgb), image.ZP, draw.Over)
}

func minMax(numbers []calc.Number) (min float64, max float64) {
	max = -math.MaxFloat64
	min = math.MaxFloat64

	for _, n := range numbers {
		f := float64(n)
		if f > max {
			max = f
		}
		if f < min {
			min = f
		}
	}
	return
}
