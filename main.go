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

	for _, token := range tokens {
		fmt.Printf("%s", token.Value)
	}
}
