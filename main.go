package main

import (
	"fmt"

	"github.com/anraku/monkey/lexer"
)

func main() {
	input := "if"

	l := lexer.New(input)

	tok := l.NextToken()
	fmt.Printf("%#v\n", tok)

}
