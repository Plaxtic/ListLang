package parser

import (
	"bufio"
	"errors"
	"exx/assemble"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type ListVars struct {
	Lists      map[string]uint
	StackSpace uint
}

type Parser struct {
	Source FileIter
	Lists  ListVars
	Body   string
}

type FileIter struct {
	Line    string
	Scanner *bufio.Scanner
	Fd      *os.File
}

func NewFileIter(filePath string) FileIter {
	file, err := os.Open(filePath)
	HandleErr(err)
	scanner := bufio.NewScanner(file)

	iter := FileIter{
		Line:    scanner.Text(),
		Fd:      file,
		Scanner: scanner,
	}
	return iter
}

func (it *FileIter) Next() bool {

	// increment scanner
	if !it.Scanner.Scan() {
		it.Fd.Close()
		return false
	}

	// update line
	it.Line = it.Scanner.Text()
	return true
}

func NewParse(filePath string) Parser {
	lists := ListVars{
		Lists:      make(map[string]uint),
		StackSpace: QWordSize * 2,
	}
	parser := Parser{
		Source: NewFileIter(filePath),
		Lists:  lists,
		Body:   "",
	}
	return parser
}

func IsValidVarName(name string) bool {
	if unicode.IsDigit(rune(name[0])) {
		return false
	}
	for _, l := range name[1:] {
		if !(unicode.IsLetter(l) || unicode.IsDigit(l)) {
			return false
		}
	}
	return true
}

func (p *Parser) NewList(name string, list string) (string, error) {
	var asm string
	if !IsValidVarName(name) {
		return asm, errors.New("Bad variable name")
	}

	asm = fmt.Sprintf(ListInit, p.Lists.StackSpace, name)
	if list != "[]" {
		asm += fmt.Sprintf(Addr2Arg1, p.Lists.StackSpace)

		list = strings.ReplaceAll(list, "[", "")
		list = strings.ReplaceAll(list, "]", "")
		items := strings.Split(list, ",")

		for _, i := range items {
			num, err := strconv.Atoi(i)
			HandleErr(err)
			asm += fmt.Sprintf(AppendList, num)
		}
	}

	p.Lists.Lists[name] = p.Lists.StackSpace
	p.Lists.StackSpace += QWordSize

	return asm, nil
}

func InKeywords(name string) bool {
	for _, k := range keyWords {
		if name == k {
			return true
		}
	}
	return false
}

func (p *Parser) InLists(name string) bool {
	if _, ok := p.Lists.Lists[name]; ok {
		return true
	} else {
		return false
	}
}

func printList(addr uint) (asm string) {
	asm = fmt.Sprintf(Addr2Arg1, addr)
	asm += PrintList
	return asm
}

func (p *Parser) LineToAsm() (string, error) {
	var asm string
	unDef := errors.New("Undefined name")

	items := strings.Split(p.Source.Line, " ")
	if len(items) < 1 || items[0] == "" {
		return asm, nil
	}

	if InKeywords(items[0]) {
		if len(items) < 2 || !p.InLists(items[1]) {
			return asm, unDef
		}
		asm = printList(p.Lists.Lists[items[1]])
		return asm, nil
	} else if p.InLists(items[0]) {
		// TODO
		return asm, nil

	} else if len(items) < 3 || items[1] != "=" {
		return asm, unDef
	} else {
		asm, err := p.NewList(items[0], strings.Join(items[2:], ""))
		return asm, err
	}
}

func (p *Parser) freeLists() string {
	var asm string
	for list := range p.Lists.Lists {
		asm += fmt.Sprintf(Addr2Arg1, p.Lists.Lists[list])
		asm += FreeLists
	}
	return asm
}

func (p *Parser) ParseFile() {
	for p.Source.Next() {
		asm, err := p.LineToAsm()
		HandleErr(err)
		p.Body += asm
	}
	p.Body += p.freeLists()
}

func (p *Parser) Compile(outFile string) {
	asm := fmt.Sprintf(Prelude, p.Lists.StackSpace)
	asm += p.Body
	asm += End
	assemble.SaveAndAssemble(outFile, asm)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
