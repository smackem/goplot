package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	port = 9090
)

func main() {
	ep := Start(port)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Running on port %d. Press Enter to quit...", port)
	reader.ReadString('\n')

	err := ep.Close()
	log.Print(err)
}
