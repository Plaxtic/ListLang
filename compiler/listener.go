package compiler

import (
	"fmt"
	"log"
	"strconv"
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

func (l *ListListener) topOfStack() int { return l.IntStack[len(l.IntStack)-1] }
func (l *ListListener) decStack()       { l.IntStack = l.IntStack[:len(l.IntStack)-1] }
func (l *ListListener) push(i int)      { l.IntStack = append(l.IntStack, i) }

func (l *ListListener) pop() (ret int) {
	if len(l.IntStack) < 1 {
		log.Panic("Cannot pop: Stack Empty")
	}
	ret = l.topOfStack()
	l.decStack()

	return ret
}

func (l *ListListener) inStack(name string) bool {
	if _, ok := l.currFunk().VarStack[name]; ok {
		return true
	} else {
		return false
	}
}

func (l *ListListener) inArgs(name string) bool {
	if _, ok := l.currFunk().Args[name]; ok {
		return true
	} else {
		return false
	}
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

func (l *ListListener) isVar(name string) bool {
	if l.inStack(name) || l.inArgs(name) {
		return true
	} else {
		return false
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

func listLen(listStr string) int {
	return strings.Count(listStr, ",")
}

func isList(listStr string) bool {
	if len(listStr) > 1 &&
		listStr[0] == '[' &&
		listStr[len(listStr)-1] == ']' {
		return true
	}
	return false
}

func (l *ListListener) appendIntStack(numItems int) {
	if numItems < 0 {
		return
	}
	pop := l.pop()
	l.appendIntStack(numItems - 1)
	l.addToFunkBody(
		fmt.Sprintf(AppendList, pop),
	)
}

func (l *ListListener) fillListVar(name, listStr string) {
	if listStr != "[]" {
		l.addToFunkBody(
			fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[name]),
		)
		l.appendIntStack(listLen(listStr))
		l.addToFunkBody(
			fmt.Sprintf(SaveRet, l.currFunk().VarStack[name]),
		)
	}
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

func isIdent(expr string) bool {
	for _, l := range expr {
		if !unicode.IsLetter(l) {
			return false
		}
	}
	return true
}

func (l *ListListener) ExitMulDivMod(c *parser.MulDivModContext) {
	opType := c.GetOp().GetTokenType()
	validateCalcType(c.GetOp().GetText(), c.GetOp().GetTokenType())

	l.Calc(
		c.Expr(0).GetText(),
		c.Expr(1).GetText(),
		opType,
	)
}

func (l *ListListener) ExitSubAdd(c *parser.SubAddContext) {
	opType := c.GetOp().GetTokenType()
	validateCalcType(c.GetOp().GetText(), opType)

	l.Calc(
		c.Expr(0).GetText(),
		c.Expr(1).GetText(),
		opType,
	)
}

func validateCalcType(op string, opType int) {
	switch opType {
	case parser.ListParserMUL:
	case parser.ListParserDIV:
	case parser.ListParserADD:
	case parser.ListParserSUB:
	default:
		log.Panicf("Unexpected op: %s\n", op)
	}
}

func (l *ListListener) Calc(left, right string, opType int) {

	isIdentL, isIdentR := isIdent(left), isIdent(right)
	switch true {
	case isIdentL && isIdentR:
		if !l.isVar(left) {
			// TODO
		} else if !l.isVar(right) {
			// TODO
		} else {
			l.bothVarMulDiv(left, right, opType)
		}
	case isIdentL:
		if !l.isVar(left) {
			fmt.Println(left, right)
			// TODO
		} else {
			l.leftVarMulDiv(left, opType)
		}
	case isIdentR:
		if !l.isVar(right) {
			// TODO
		} else {
			l.leftVarMulDiv(right, opType)
		}
	default:
		l.constMulDiv(opType)
	}
}

func (l *ListListener) bothVarMulDiv(left, right string, opType int) {

	var asmMacro string
	switch opType {
	case parser.ListParserMUL:
		asmMacro = MulTwoVar
	case parser.ListParserDIV:
		asmMacro = DivTwoVar
	case parser.ListParserADD:
		asmMacro = AddTwoVar
	case parser.ListParserSUB:
		asmMacro = SubTwoVar
	}
	l.addToFunkBody(
		fmt.Sprintf(asmMacro,
			l.getVar(left),
			l.getVar(right),
		),
	)
}

func (l *ListListener) leftVarMulDiv(left string, opType int) {

	var asmMacro string
	switch opType {
	case parser.ListParserMUL:
		asmMacro = MulOneVarLeft
	case parser.ListParserDIV:
		asmMacro = DivOneVarLeft
	case parser.ListParserADD:
		asmMacro = AddOneVarLeft
	case parser.ListParserSUB:
		asmMacro = SubOneVarLeft
	}
	l.addToFunkBody(
		fmt.Sprintf(asmMacro,
			l.currFunk().VarStack[left],
			l.pop(),
		),
	)
}

func (l *ListListener) rightVarMulDiv(right string, opType int) {

	var asmMacro string
	switch opType {
	case parser.ListParserMUL:
		asmMacro = MulOneVarRight
	case parser.ListParserDIV:
		asmMacro = DivOneVarRight
	case parser.ListParserADD:
		asmMacro = AddOneVarRight
	case parser.ListParserSUB:
		asmMacro = SubOneVarRight
	}
	l.addToFunkBody(
		fmt.Sprintf(asmMacro,
			l.pop(),
			l.currFunk().VarStack[right],
		),
	)
}

func (l *ListListener) constMulDiv(opType int) {
	right, left := l.pop(), l.pop()

	switch opType {
	case parser.ListParserMUL:
		l.push(left * right)
	case parser.ListParserDIV:
		l.push(left / right)
	case parser.ListParserADD:
		l.push(left + right)
	case parser.ListParserSUB:
		l.push(left - right)
	}
}

func (l *ListListener) EnterNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	HandleErr(err)
	l.push(i)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
