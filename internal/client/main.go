package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var steps int
	var minY float64
	var maxY float64

	flag.IntVar(&steps, "steps", 1000, "Horizontal resolution.")
	flag.Float64Var(&minY, "miny", 0, "Lower Y bound. Defaults to minimum value in function results.")
	flag.Float64Var(&maxY, "maxy", 0, "Upper Y bound. Defaults to maximum value in function results.")

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Printf("%s [OPTIONS] URL FUNCTION\nOPTIONS:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	urlstr := args[0]
	fsrc := url.QueryEscape(args[1])

	urlstr = fmt.Sprintf("%s?f=%s&steps=%d&miny=%g&maxy=%g", urlstr, fsrc, steps, minY, maxY)
	fmt.Fprintf(os.Stderr, "> GET %s\n", urlstr)

	resp, err := http.Get(urlstr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintf(os.Stderr, "< %s %s\n", resp.Proto, resp.Status)
	for key, val := range resp.Header {
		fmt.Fprintf(os.Stderr, "< %s: %s\n", key, val)
	}

	count, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintf(os.Stderr, "%d bytes read\n", count)
}
