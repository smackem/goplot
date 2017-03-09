package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

// EndPoint represents an HTTP listener bound to a specific port.
type EndPoint interface {
	// Closes the endpoint.
	Close() error
}

// Start binds the specified port and starts the listener.
func Start(port int) EndPoint {
	http.HandleFunc("/", handler)

	s := &httpServer{
		srv: http.Server{
			Addr:           fmt.Sprintf(":%d", port),
			Handler:        nil, // use DefaultServeMux
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		closedSignal: make(chan error),
	}

	go func() {
		err := s.srv.ListenAndServe()
		log.Print(err)
		s.closedSignal <- err
	}()

	return s
}

type httpServer struct {
	srv          http.Server
	closedSignal chan error
}

func (s *httpServer) Close() error {
	s.srv.Close()
	return <-s.closedSignal
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", html.EscapeString(r.URL.Path[1:]))
}
