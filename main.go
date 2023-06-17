package main

import (
	"fmt"
	"ink/lexer"
)

func main() {
	code := `defun()`

	lexer := lexer.NewLexer(code)

	tokens := lexer.MakeTokens()
	for _, token := range tokens {
		fmt.Printf("Kind: %s, Value: %s\n", token.Kind, token.Value)
	}
}
