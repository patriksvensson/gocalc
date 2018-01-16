package main

import (
	"fmt"
	"github.com/patriksvensson/gocalc/parsing"
)

func main() {

	tokens, err := parsing.Tokenize("12 + (4 - 5)")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", parsing.Parse(&tokens))
}
