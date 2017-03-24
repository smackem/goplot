package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("USAGE: go run main.go URL FUNCTION")
		return
	}

	urlstr := os.Args[1]
	fsrc := url.QueryEscape(os.Args[2])

	if strings.Contains(urlstr, "?") {
		urlstr += "&"
	} else {
		urlstr += "?"
	}

	urlstr += "f=" + fsrc
	fmt.Fprintf(os.Stderr, "> GET %s\n", urlstr)

	resp, err := http.Get(urlstr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stderr, "< %s %s\n", resp.Proto, resp.Status)
	for key, val := range resp.Header {
		fmt.Fprintf(os.Stderr, "< %s: %s\n", key, val)
	}

	os.Stdout.Write(b)
}
