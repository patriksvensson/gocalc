package main

import (
	"fmt"
	"github.com/patriksvensson/gocalc/parser"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("usage: gocalc <expression>")
		return
	}

	exp := strings.Join(args, " ")
	res, err := parser.Evaluate(exp)
	if err != nil {
		fmt.Printf("Ooops! %s\n", err)
		return
	}

	fmt.Printf("%d", res)
}
