package main

import "fmt"
import "net/http"

const (
	port = 9090
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Printf("Listening on port %d\n", port)

	pattern := fmt.Sprintf(":%d", port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(pattern, nil)
}
