package main

import (
	"fmt"
	"test/compiler"
	"test/output"
	"test/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {

	out := output.NewOutput("test.list")
	input := out.LoadFile()

	is := antlr.NewInputStream(input)

	lexer := parser.NewListLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewListParser(stream)

	listener := compiler.NewListListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	for id, addr := range listener.Stack {
		fmt.Printf("%-10s @ rbp+%#x\n", id, addr)
	}
	output := listener.AsmPrelude + listener.AsmBody + listener.AsmEpilogue
	//	fmt.Println(output)
	out.WriteAsm(output)
}
