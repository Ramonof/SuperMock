package main

import (
	"SuperStub/internal/goroovy"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}

	tokens := make([]*goroovy.Tokenized, 0)
	newLexer := goroovy.NewLexer(file)
	for {
		pos, tok, lit := newLexer.Lex()
		if tok == goroovy.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
		tokens = append(tokens, &goroovy.Tokenized{Line: pos.Line, Col: pos.Column, Token: tok, Lit: lit})
	}

	newParser := goroovy.NewParser(tokens)
	newParser.AddVariable("res1", "{\"test\":\"body\"}")
	newParser.AddVariable("request.body.id", "1")
	res, err := newParser.ParseTokens()
	if err != nil {
		panic(err)
	}

	fmt.Printf(res)
}
