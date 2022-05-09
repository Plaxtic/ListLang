package compiler

import (
	"fmt"
	"log"
	"strings"
	"test/parser"
	"unicode"
)

type ListListener struct {
	*parser.BaseListListener

	IntStack  []int
	FunkStack []string

	Funks map[string]*StackFrame
}

func NewListListener() *ListListener {
	newList := &ListListener{
		Funks: make(map[string]*StackFrame),
	}
	newList.initFunk("main")

	return newList
}

func listLen(listStr string) int {
	return strings.Count(listStr, ",")
}

func isList(listStr string) bool {
	return len(listStr) > 1 &&
		listStr[0] == '[' &&
		listStr[len(listStr)-1] == ']'
}

func isIdent(expr string) bool {
	for _, l := range expr {
		if !unicode.IsLetter(l) {
			return false
		}
	}
	return true
}

func inAddrMap(name string, addrMap map[string]uint) bool {
	_, ok := addrMap[name]
	return ok
}

func (l *ListListener) inStack(name string) bool {
	return inAddrMap(name, l.currFunk().VarStack)
}

func (l *ListListener) inArgs(name string) bool {
	return inAddrMap(name, l.currFunk().Args)
}

func (l *ListListener) isVar(name string) bool {
	return l.inStack(name) || l.inArgs(name)
}

func (l *ListListener) getVar(name string) uint {
	if l.inStack(name) {
		return l.currFunk().VarStack[name]
	} else if l.inArgs(name) {
		return l.currFunk().Args[name]
	} else {
		log.Panicf("Undeclared identifier '%s'", name)
		return 0
	}
}

func (l *ListListener) addVarToStack(name string) {
	if l.inStack(name) {
		// reassign, free if orphan pointer
		// TODO
	} else {
		l.currFunk().VarStack[name] = l.currFunk().StackSize
		l.currFunk().StackSize += QWordSize
	}
}

func (l *ListListener) ExitStart(c *parser.StartContext) {
	main := l.Funks["main"]
	main.AsmPrelude = fmt.Sprintf(Prelude, l.currFunk().StackSize)

	for _, a := range main.VarStack {
		main.AsmEpilogue += fmt.Sprintf(Addr2Arg1, a)
		main.AsmEpilogue += FreeLists
	}
	main.AsmEpilogue += End
}

func (l *ListListener) GetAsm() (fullAsm string) {
	for _, asm := range l.Funks {
		fullAsm += asm.AsmPrelude
		fullAsm += asm.AsmBody
		fullAsm += asm.AsmEpilogue
	}
	return fullAsm
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
