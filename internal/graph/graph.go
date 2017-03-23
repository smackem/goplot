package graph

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/smackem/goplot/internal/calc"
)

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

func PlotPng(xs []calc.Number, ys []calc.Number) image.Image {
	if len(xs) <= 0 || len(xs) != len(ys) {
		panic("xs and ys must be of equal length, both greater than zero")
	}

	bounds := image.Rect(0, 0, 800, 600)

	dest := image.NewRGBA(bounds)
	gc := draw2dimg.NewGraphicContext(dest)

	// fill with white
	draw.Draw(dest, bounds, image.White, image.ZP, draw.Over)

	gc.SetStrokeColor(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
	gc.SetLineWidth(1)

	// transform to center 0,0
	minX, maxX := -float64(bounds.Dx())/2, float64(bounds.Dx())/2
	minY, maxY := -float64(bounds.Dy())/2, float64(bounds.Dy())/2
	gc.Translate(maxX+0.5, maxY+0.5)

	// draw x and y axis
	gc.MoveTo(minX, 0)
	gc.LineTo(maxX, 0)
	gc.Stroke()

	gc.MoveTo(0, minY)
	gc.LineTo(0, maxY)
	gc.Stroke()

	// draw curve
	gc.SetLineWidth(3)
	gc.SetStrokeColor(color.RGBA{0, 0, 0, 0xff})
	gc.MoveTo(float64(xs[0]), float64(ys[0]))
	for i := 1; i < len(xs); i++ {
		x, y := float64(xs[i]), -float64(ys[i])
		gc.LineTo(x, y)
	}
	gc.Stroke()

	return dest
}

func minmax(slice []calc.Number) (min float64, max float64) {
	max = -math.MaxFloat64
	min = math.MaxFloat64
	for _, n := range slice {
		f := float64(n)
		if f < min {
			min = f
		}
		if f > max {
			max = f
		}
	}
	return
}
