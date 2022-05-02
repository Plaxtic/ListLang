package cli

import (
	"log"
	"os"
	"test/compiler"
	"test/output"
	"test/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Cli struct {
	Input    string
	Listener *compiler.ListListener
	Asm      string
}

func (c *Cli) compile() {

	c.Listener = compiler.NewListListener()

	p := parser.NewListParser(
		antlr.NewCommonTokenStream(
			parser.NewListLexer(
				antlr.NewInputStream(
					c.Input,
				),
			),
			antlr.TokenDefaultChannel,
		),
	)
	antlr.ParseTreeWalkerDefault.Walk(c.Listener, p.Start())

	c.Asm += c.Listener.AsmPrelude
	c.Asm += c.Listener.AsmBody
	c.Asm += c.Listener.AsmEpilogue
}

func Run() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <source>\n", os.Args[0])
	}

	out := output.NewOutput(os.Args[1])
	cli := Cli{
		Input: out.LoadFile(),
	}
	cli.compile()

	//	for id, addr := range cli.Listener.Stack {
	//		fmt.Printf("%-10s @ rbp+%#x\n", id, addr)
	//	}

	//	fmt.Println(cli.Asm)
	out.WriteAsm(cli.Asm)
}
