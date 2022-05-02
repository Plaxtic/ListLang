package compiler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"test/parser"
)

type ListListener struct {
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

func NewListListener() *ListListener {
	newList := &ListListener{
		Stack:        make(map[string]uint64),
		StackPointer: RetAndRpb,
	}
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
	if _, ok := l.Stack[name]; ok {
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
		l.Stack[name] = uint64(l.StackPointer)
		l.StackPointer += QWordSize
	}
}

func (l *ListListener) initializeList(name string) {
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

func (l *ListListener) appendIntStack(numItems int) {
	if numItems < 0 {
		return
	}
	pop := l.pop()
	l.appendIntStack(numItems - 1)
	l.AsmBody += fmt.Sprintf(AppendList, pop)
}

func (l *ListListener) fillList(name string, list parser.IListTypeContext) {
	if list.GetText() == "[]" {
		return
	} else {
		l.AsmBody += fmt.Sprintf(Addr2Arg1, l.Stack[name])
		l.appendIntStack(listLen(list.GetText()))
		l.AsmBody += fmt.Sprintf(SaveRet, l.Stack[name])
	}
}

func (l *ListListener) ExitAssign(c *parser.AssignContext) {
	fmt.Println("ExitAssign")

	name := c.IDENT().GetText()
	l.addVarToStack(name)
	l.initializeList(name)
	l.fillList(name, c.ListType())
}

func (l *ListListener) EnterDeclaration(c *parser.DeclarationContext) {
	fmt.Println("EnterDeclaration")
}

func (l *ListListener) EnterSendRight(c *parser.SendRightContext) {
	fmt.Println("EnterSendRight")
}

func (l *ListListener) printList(ident string) {
	if !l.inStack(ident) {
		log.Panicf("Unknown identifier '%s'\n", ident)
	} else {
		l.AsmBody += fmt.Sprintf(Addr2Arg1, l.Stack[ident])
		l.AsmBody += PrintList
	}
}

func (l *ListListener) sendListL(send string, recv string) {
	if !l.inStack(send) {
		log.Panicf("Unknown identifier '%s'\n", send)
	} else {
		l.AsmBody += fmt.Sprintf(Addr2Arg1, l.Stack[send])
		l.AsmBody += fmt.Sprintf(Int2Arg2, fileDescriptors[recv])
		l.AsmBody += SendListL
	}
}

func (l *ListListener) ExitSendLeft(c *parser.SendLeftContext) {
	fmt.Println("ExitSendLeft")

	recv, send := c.IDENT(0).GetText(), c.IDENT(1).GetText()
	if isConstFd(recv) {
		switch recv {
		case "stdout":
			l.sendListL(send, recv)
		default:
			fmt.Printf("%s not implemented\n", recv)
		}
	}
}

func (l *ListListener) sendListR(send string, recv string) {
	if !l.inStack(send) {
		log.Panicf("Unknown identifier '%s'\n", send)
	} else {
		l.AsmBody += fmt.Sprintf(Addr2Arg1, l.Stack[send])
		l.AsmBody += fmt.Sprintf(Int2Arg2, fileDescriptors[recv])
		l.AsmBody += SendListR
	}
}

func (l *ListListener) ExitSendRight(c *parser.SendRightContext) {
	fmt.Println("ExitSendRight")

	send, recv := c.IDENT(0).GetText(), c.IDENT(1).GetText()
	if isConstFd(recv) {
		switch recv {
		case "stdout":
			l.sendListR(send, recv)
		default:
			fmt.Printf("%s not implemented\n", recv)
		}
	}
}

func (l *ListListener) EnterLine(c *parser.LineContext) {
	fmt.Println("EnterLine")
}

func (l *ListListener) ExitStart(c *parser.StartContext) {
	l.AsmPrelude = fmt.Sprintf(Prelude, l.StackPointer)

	for _, a := range l.Stack {
		l.AsmEpilogue += fmt.Sprintf(Addr2Arg1, a)
		l.AsmEpilogue += FreeLists
	}
	l.AsmEpilogue += End
}

// ONLY WORKS FOR CONSTANT VALUES
func (l *ListListener) ExitMulDiv(c *parser.MulDivContext) {
	//	fmt.Println("EnterCalc")
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.ListParserMUL:
		l.push(left * right)
	case parser.ListParserDIV:
		l.push(left / right)
	default:
		log.Panicf("Unexpected op: %s\n", c.GetOp().GetText())
	}
}

func (l *ListListener) ExitSubAdd(c *parser.SubAddContext) {
	//	fmt.Println("EnterCalc")
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.ListParserADD:
		l.push(left + right)
	case parser.ListParserSUB:
		l.push(left - right)
	default:
		log.Panicf("Unexpected op: %s\n", c.GetOp().GetText())
	}
}

func (l *ListListener) EnterNumber(c *parser.NumberContext) {
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

//func (l *ListListener) EnterVariable(ctx *parser.VariableContext) {
//	fmt.Println("EnterVariable")
//}

func (l *ListListener) ExitExprList(c *parser.ExprListContext) {
	fmt.Println("ExitExprList")
	for _, e := range c.AllExpr() {
		fmt.Printf("e: %s\n", e.GetText())
	}
}

func (l *ListListener) ExitExpr(c *parser.ExprContext) {
	fmt.Println("ExitExpr")
}
