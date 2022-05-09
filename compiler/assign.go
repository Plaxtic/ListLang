package compiler

import (
	"fmt"
	"log"
	"test/parser"
)

func (l *ListListener) appendIntStack(numItems int) {
	if numItems < 0 {
		return
	}

	popInt := l.popInt()
	l.appendIntStack(numItems - 1)
	l.addToFunkBody(
		fmt.Sprintf(AppendList, popInt),
	)
}

func (l *ListListener) fillListVar(name, listStr string) {
	if listStr == "[]" {
		return
	}

	l.addToFunkBody(
		fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[name]),
	)
	l.appendIntStack(listLen(listStr))
	l.addToFunkBody(
		fmt.Sprintf(SaveRet, l.currFunk().VarStack[name]),
	)
}

// leaves reference in rax and rdi
func (l *ListListener) fillRawList(listStr string) {
	l.addToFunkBody(ListInit + Ret2Arg1)

	if listStr != "[]" {
		l.appendIntStack(listLen(listStr))
	}
}

func (l *ListListener) ExitAssign(c *parser.AssignContext) {
	name := c.IDENT().GetText()

	l.addVarToStack(name)
	l.addToFunkBody(ListInit)
	l.addToFunkBody(
		fmt.Sprintf(SaveRet, l.currFunk().VarStack[name]),
	)

	l.fillListVar(name, c.ListType().GetText())
}

func (l *ListListener) printList(ident string) {
	if !l.inStack(ident) {
		log.Panicf("Unknown identifier '%s'\n", ident)
	} else {
		l.addToFunkBody(
			fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[ident]) +
				PrintList,
		)
	}
}
