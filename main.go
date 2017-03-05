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
	calculator := calc.Calculator{}
	calculator.Evaluate("1+1")
	tokens := calc.Lex(" ( ) + - * ")
	log.Print(tokens)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Running on port %d. Press Enter to quit...", port)
	reader.ReadString('\n')

	err := ep.Close()
	log.Print(err)
}
