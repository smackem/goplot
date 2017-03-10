package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/smackem/goplot/internal/calc"
	"github.com/smackem/goplot/internal/graph"
)

func registerAPI() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/eval", evalHandler)
	http.HandleFunc("/draw", drawHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}
	body := string(bytes)
	fmt.Fprintf(w, "Got '%s'!\n", html.EscapeString(body))
}

func drawHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/png")
	graph.DrawPng(w)
}

func evalHandler(w http.ResponseWriter, r *http.Request) {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	steps, err := strconv.Atoi(uri.Query().Get("steps"))
	if err != nil {
		steps = 10
	}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, err := calc.Parse(string(bytes))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	xs, ys := f.Eval(steps)

	w.Header().Set("Content-type", "application/json; charset=utf-8")
	result := struct {
		Xs []calc.Number
		Ys []calc.Number
	}{Xs: xs, Ys: ys}
	enc := json.NewEncoder(w)
	enc.Encode(result)
}
