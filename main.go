package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = 9090
)

func main() {
	registerAPI()

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	fmt.Printf("Running on port %d. Press Ctrl+C to quit...", port)
	err := srv.ListenAndServe()
	log.Print(err)
}
