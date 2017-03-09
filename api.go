package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smackem/goplot/internal/calc"
)

func registerAPI() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/eval", evalHandler)
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

func evalHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}
	body := string(bytes)
	f, err := calc.Parse(body)
	if err != nil {
		log.Printf("Error parsing function: %v", err)
		return
	}
	w.Header().Set("Content-type", "text/plain; encoding=utf-8")
	xs, ys := f.Eval(10)
	fmt.Fprintf(w, "xs: %v, ys: %v\n", xs, ys)
}
