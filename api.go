package main

import (
	"encoding/xml"
	"fmt"

	"github.com/smackem/goplot/internal/calc"
	"github.com/smackem/goplot/internal/goobar"
	"github.com/smackem/goplot/internal/graph"
)

func registerAPI() {
	goobar.RegisterAction("/", rootAction)
	goobar.RegisterAction("/eval", evalAction)
	goobar.RegisterAction("/draw", drawAction)
	goobar.RegisterAction("/plot", plotAction)
	goobar.RegisterFileServer("/public/", "resource/public")
	goobar.RegisterViewPath("resource/view")
}

func rootAction(x goobar.Exchange) goobar.Responder {
	text := fmt.Sprintf("Got '%d'!\n", x.MustGetID())
	return goobar.XML(struct {
		XMLName xml.Name `xml:"v"`
		Msg     string
		Number  int
	}{Msg: text, Number: 123})
	//return goobar.PlainText(text)
}

func drawAction(x goobar.Exchange) goobar.Responder {
	return goobar.ImagePNG(graph.DrawPng())
}

func plotAction(x goobar.Exchange) goobar.Responder {
	fsrc := x.MustGetString("f")
	f, err := calc.Parse(fsrc)
	if err != nil {
		return goobar.Error(400, err.Error())
	}
	steps := x.GetIntOrDefault("steps", 10)
	minY := x.GetFloatOrDefault("miny", 0)
	maxY := x.GetFloatOrDefault("maxy", 0)

	xs, ys := f.Eval(steps)
	return goobar.ImagePNG(graph.PlotPng(xs, ys, minY, maxY))
}

func evalAction(x goobar.Exchange) goobar.Responder {
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
