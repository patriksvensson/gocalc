package main

import (
	"fmt"
	"github.com/patriksvensson/gocalc/parser"
)

func main() {
	res, err := parser.Evaluate("1 + (1 - 3) + 7 - (1)")
	if err != nil {
		fmt.Printf("Ooops! %s", err)
		return
	}

	fmt.Printf("%d", res)
}
