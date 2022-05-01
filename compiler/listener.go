package compiler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"test/parser"
)

type listListener struct {
	*parser.BaseListListener

	IntStack []int

	Stack        map[string]uint64
	StackPointer uint64

	AsmPrelude  string
	AsmBody     string
	AsmEpilogue string
}

const (
	RetAndRpb = QWordSize * 2
)

func NewListListener() *listListener {
	newList := &listListener{
		Stack:        make(map[string]uint64),
		StackPointer: RetAndRpb,
	}
	return newList
}

func (l *listListener) topOfStack() int { return l.IntStack[len(l.IntStack)-1] }
func (l *listListener) decStack()       { l.IntStack = l.IntStack[:len(l.IntStack)-1] }
func (l *listListener) push(i int)      { l.IntStack = append(l.IntStack, i) }

func (l *listListener) pop() (ret int) {
	if len(l.IntStack) < 1 {
		log.Panic("Cannot pop: Stack Empty")
	}
	ret = l.topOfStack()
	l.decStack()

	return ret
}

func (l *listListener) addVarToStack(name string) {
	if _, ok := l.Stack[name]; ok {
		// reassign, free if orphan pointer
		// TODO
	} else {
		l.Stack[name] = uint64(l.StackPointer)
		l.StackPointer += QWordSize
	}
}

func (l *listListener) initializeList(name string) {
	l.AsmBody += fmt.Sprintf(ListInit, l.Stack[name], name)
}

// IMPROVE / REMOVE
func parseList(listStr string) (list []uint64) {
	listStr = strings.Trim(listStr, "[]")
	items := strings.Split(listStr, ",")

	for _, item := range items {
		i, err := strconv.ParseUint(item, 10, 64)
		HandleErr(err)
		list = append(list, i)
	}
	return list
}

func listLen(listStr string) int {
	return strings.Count(listStr, ",")
}

func (l *listListener) appendIntStack(numItems int) {
	if numItems < 0 {
		return
	}
	pop := l.pop()
	l.appendIntStack(numItems - 1)
	l.AsmBody += fmt.Sprintf(AppendList, pop)
}

func (l *listListener) fillList(name string, list parser.IListTypeContext) {
	if list.GetText() == "[]" {
		return
	} else {
		l.AsmBody += fmt.Sprintf(Addr2Arg1, l.Stack[name])
		l.appendIntStack(listLen(list.GetText()))
		l.AsmBody += fmt.Sprintf(SaveRet, l.Stack[name])
	}
}

func (l *listListener) ExitAssign(c *parser.AssignContext) {
	fmt.Println("ExitAssign")
	name := c.IDENT().GetText()
	l.addVarToStack(name)
	l.initializeList(name)
	l.fillList(name, c.ListType())
}

func (l *listListener) ExitStart(c *parser.StartContext) {

	l.AsmPrelude = fmt.Sprintf(Prelude, l.StackPointer)

	for _, a := range l.Stack {
		l.AsmEpilogue += fmt.Sprintf(Addr2Arg1, a)
		l.AsmEpilogue += FreeLists
	}
	l.AsmEpilogue += End
}

// ONLY WORKS FOR CONSTANT VALUES
func (l *listListener) ExitCalc(c *parser.CalcContext) {
	fmt.Println("EnterCalc")
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.ListParserMUL:
		l.push(left * right)
	case parser.ListParserDIV:
		l.push(left / right)
	case parser.ListParserADD:
		l.push(left + right)
	case parser.ListParserSUB:
		l.push(left - right)
	default:
		log.Panicf("Unexpected op: %s\n", c.GetOp().GetText())
	}
}

func (l *listListener) EnterNumber(c *parser.NumberContext) {
	fmt.Println("EnterNumber")
	i, err := strconv.Atoi(c.GetText())
	HandleErr(err)
	l.push(i)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (l *listListener) EnterVariable(ctx *parser.VariableContext) {
	fmt.Println("EnterVariable")
}

func (l *listListener) ExitExprList(c *parser.ExprListContext) {
	fmt.Println("ExitExprList")
	for _, e := range c.AllExpr() {
		fmt.Printf("e: %s\n", e.GetText())
	}
}
func (l *listListener) ExitExpr(c *parser.ExprContext) {
	fmt.Println("ExitExpr")
}
