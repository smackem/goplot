package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/smackem/goplot/internal/calc"
	"github.com/smackem/goplot/internal/goobar"
	"github.com/smackem/goplot/internal/graph"
)

func registerHTTP() {
	http.Handle("/api/eval", goobar.Get(getEval))

	http.Handle("/plot/img", goobar.Get(getPlotImg))
	http.Handle("/plot/svg", goobar.Get(getPlotSvg))

	http.Handle("/", goobar.Get(getIndex))
	http.Handle("/svg", goobar.Get(getSvgIndex))

	http.Handle("/pub/", http.FileServer(http.Dir("resource")))
}

func getPlotImg(x *goobar.Exchange) goobar.Responder {
	fsrc := x.MustGetString("f")
	f, err := calc.Parse(fsrc)
	if err != nil {
		return goobar.Error(400, err.Error())
	}
	steps := x.GetIntOrDefault("steps", 100)
	minY := x.GetFloatOrDefault("miny", 0)
	maxY := x.GetFloatOrDefault("maxy", 0)

	xs, ys := f.Eval(steps)
	return goobar.ImagePNG(graph.PlotPng(xs, ys, minY, maxY))
}

func getPlotSvg(x *goobar.Exchange) goobar.Responder {
	fsrc := x.MustGetString("f")
	f, err := calc.Parse(fsrc)
	if err != nil {
		return goobar.Error(400, err.Error())
	}
	steps := x.GetIntOrDefault("steps", 100)
	minY := x.GetFloatOrDefault("miny", 0)
	maxY := x.GetFloatOrDefault("maxy", 0)

	xs, ys := f.Eval(steps)
	tg := graph.Transform(xs, ys, minY, maxY, 600, 400)
	pointsBuf := new(bytes.Buffer)
	const pad = 2.5

	for i, x := range tg.Xs {
		if i > 0 {
			pointsBuf.WriteRune(' ')
		}
		fmt.Fprintf(pointsBuf, "%.2f,%.2f", x+pad, tg.Ys[i]+pad)
	}
	return goobar.View("svg/plot.svg", struct {
		Points         string
		XAxisY, YAxisX float64
		Width, Height  float64
	}{
		pointsBuf.String(),
		tg.XAxisY + pad, tg.YAxisX + pad,
		tg.Width + pad*2, tg.Height + pad*2,
	})
}

func getEval(x *goobar.Exchange) goobar.Responder {
	fsrc := x.MustGetString("f")
	f, err := calc.Parse(fsrc)
	if err != nil {
		return goobar.Error(400, err.Error())
	}
	steps := x.GetIntOrDefault("steps", 10)

	xs, ys := f.Eval(steps)
	return goobar.JSON(struct {
		Xs []calc.Number
		Ys []calc.Number
	}{xs, ys})
}

func getIndex(x *goobar.Exchange) goobar.Responder {
	return goobar.View("index.html", struct {
		Message string
		Link    string
	}{"Hello", "/draw"})
}

func getSvgIndex(x *goobar.Exchange) goobar.Responder {
	return goobar.View("svg/index.html", nil)
}
