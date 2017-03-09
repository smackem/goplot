package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port = 9090
)

func main() {
	registerAPI()

	closeChannel := make(chan error)
	srv := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		closeChannel <- srv.ListenAndServe()
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Running on port %d. Press Enter to quit...", port)
	reader.ReadString('\n')

	srv.Close()
	log.Print(<-closeChannel)
}
