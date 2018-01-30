package main

import (
	"fmt"
	"github.com/patriksvensson/gocalc/parser"
	"os"
	"strings"
)

func main() {

	parser.Evaluate("1+2-3+4")

	// We need at least one argument.
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("usage: gocalc <expression>")
		return
	}

	// Get the full expression by joining
	// all arguments together.
	exp := strings.Join(args, " ")

	// Evaluate the expression.
	res, err := parser.Evaluate(exp)
	if err != nil {
		fmt.Printf("Ooops! %s\n", err)
		return
	}

	// Print the result.
	fmt.Printf("%d", res)
}
