package main

import (
	"exx/parser"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Panicf("Usage: %s <source>\n", os.Args[0])
	}
	parser := parser.NewParse(os.Args[1])
	parser.ParseFile()
	parser.Compile(os.Args[1])
}
