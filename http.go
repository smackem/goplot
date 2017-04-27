package main

import (
	"net/http"

	"github.com/smackem/goplot/internal/calc"
	"github.com/smackem/goplot/internal/goobar"
	"github.com/smackem/goplot/internal/graph"
)

func registerHTTP() {
	http.Handle("/api/eval", goobar.Get(getEval))

	http.Handle("/plot/img", goobar.Get(getPlotImg))
	http.Handle("/plot/svg", goobar.Get(getPlotImg))

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
