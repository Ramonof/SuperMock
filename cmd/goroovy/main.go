package main

import (
	"SuperStub/cmd/goroovy/lexer"
	"SuperStub/cmd/goroovy/parser"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}

	tokens := make([]*parser.Token, 0)
	newLexer := lexer.NewLexer(file)
	for {
		pos, tok, lit := newLexer.Lex()
		if tok == lexer.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
		tokens = append(tokens, &parser.Token{Line: pos.Line, Col: pos.Column, Token: tok, Lit: lit})
	}

	newParser := parser.NewParser(tokens)
	newParser.AddVariable("res1", "{\"test\":\"body\"}")
	newParser.AddVariable("request.body.id", "1")
	res, err := newParser.ParseTokens()
	if err != nil {
		panic(err)
	}

	fmt.Printf(res)
}
