package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/smackem/goplot/internal/calc"
)

const (
	port = 9090
)

func main() {
	ep := Start(port)
	calc := calc.Calculator{}
	calc.Evaluate("1+1")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Running on port %d. Press Enter to quit...", port)
	reader.ReadString('\n')

	err := ep.Close()
	log.Print(err)
}
