package compiler

import (
	"fmt"
	"strings"
	"test/parser"
)

type StackFrame struct {
	Args map[string]uint

	VarStack  map[string]uint
	StackSize uint

	AsmPrelude  string
	AsmBody     string
	AsmEpilogue string
}

func (l *ListListener) initFunk(funk string) {
	l.Funks[funk] = &StackFrame{
		Args:      make(map[string]uint),
		VarStack:  make(map[string]uint),
		StackSize: RetAndRpb,
	}
	l.pushFunk(funk)
}

func (l *ListListener) topOfFunkStack() string {
	return l.FunkStack[len(l.FunkStack)-1]
}
func (l *ListListener) decFunkStack() {
	l.FunkStack = l.FunkStack[:len(l.FunkStack)-1]
}
func (l *ListListener) pushFunk(name string) {
	l.FunkStack = append(l.FunkStack, name)
}

func (l *ListListener) addToFunkBody(asm string) {
	l.currFunk().AsmBody += asm
}

func (l *ListListener) currFunk() *StackFrame {
	return l.Funks[l.topOfFunkStack()]
}

func (l *ListListener) addArg(arg string) {
	l.currFunk().Args[arg] = l.currFunk().StackSize + RetAndRpb
	l.currFunk().StackSize += QWordSize
}

func (l *ListListener) ExitArgs(c *parser.ArgsContext) {
	args := strings.Split(c.GetText(), ",")

	for _, arg := range args {
		l.addArg(arg)
	}
}

func (l *ListListener) EnterFuncDec(c *parser.FuncDecContext) {
	fmt.Println("EnterFuncDec")

	funk := c.IDENT().GetText()
	l.initFunk(funk)
}

func (l *ListListener) ExitFuncDec(c *parser.FuncDecContext) {
	fmt.Println("ExitFuncDec")

	curr := l.currFunk()
	funk := c.IDENT().GetText()
	curr.AsmPrelude = fmt.Sprintf(
		FnPrelude,
		funk,
		l.currFunk().StackSize,
	)

	for _, a := range curr.VarStack {
		curr.AsmEpilogue += fmt.Sprintf(Addr2Arg1, a)
		curr.AsmEpilogue += FreeLists
	}
	curr.AsmEpilogue += End

	l.decFunkStack()
}
