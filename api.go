package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/smackem/goplot/internal/calc"
	"github.com/smackem/goplot/internal/goobar"
	"github.com/smackem/goplot/internal/graph"
)

func registerAPI() {
	http.Handle("/public/", http.FileServer(http.Dir("resource")))
	http.Handle("/", goobar.Get(getRoot))
	http.Handle("/eval", goobar.Get(getEval))
	http.Handle("/draw", goobar.Get(getDraw))
	http.Handle("/plot", goobar.Get(getPlot))
	http.Handle("/index", goobar.Get(getIndex))
	http.Handle("/sub", goobar.Get(func(x *goobar.Exchange) goobar.Responder {
		return goobar.View("sub/subfile.html", nil)
	}))
}

func getRoot(x *goobar.Exchange) goobar.Responder {
	text := fmt.Sprintf("Got '%d'!\n", x.MustGetID())
	return goobar.XML(struct {
		XMLName xml.Name `xml:"v"`
		Msg     string
		Number  int
	}{Msg: text, Number: 123})
	//return goobar.PlainText(text)
}

func getDraw(x *goobar.Exchange) goobar.Responder {
	return goobar.ImagePNG(graph.DrawPng())
}

func getPlot(x *goobar.Exchange) goobar.Responder {
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
