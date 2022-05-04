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

	IntStack  []int
	FunkStack []string

	VarStack  map[string]uint64
	Funks     map[string]*StackFrame
	StackSize uint64
}

type StackFrame struct {
	AsmPrelude  string
	AsmBody     string
	AsmEpilogue string
}

func NewListListener() *ListListener {
	newList := &ListListener{
		VarStack:  make(map[string]uint64),
		Funks:     make(map[string]*StackFrame),
		StackSize: RetAndRpb,
	}
	newList.initFunk("main")

	return newList
}

func (l *ListListener) initFunk(funk string) {
	l.Funks[funk] = &StackFrame{}
	l.pushFunk("main")
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

func (l *ListListener) topOfFunkStack() string { return l.FunkStack[len(l.FunkStack)-1] }
func (l *ListListener) decFunkStack()          { l.FunkStack = l.FunkStack[:len(l.FunkStack)-1] }
func (l *ListListener) pushFunk(name string)   { l.FunkStack = append(l.FunkStack, name) }

func (l *ListListener) addToFunkBody(asm string) {
	l.Funks[l.topOfFunkStack()].AsmBody += asm
}

func (l *ListListener) inStack(name string) bool {
	if _, ok := l.VarStack[name]; ok {
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
		l.VarStack[name] = uint64(l.StackSize)
		l.StackSize += QWordSize
	}
}

func (l *ListListener) initializeList(name string) {
	l.addToFunkBody(
		fmt.Sprintf(ListInit, l.VarStack[name], name),
	)
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
	l.addToFunkBody(
		fmt.Sprintf(AppendList, pop),
	)
}

func (l *ListListener) fillList(name string, list parser.IListTypeContext) {
	if list.GetText() == "[]" {
		return
	} else {
		l.addToFunkBody(
			fmt.Sprintf(Addr2Arg1, l.VarStack[name]),
		)
		l.appendIntStack(listLen(list.GetText()))
		l.addToFunkBody(
			fmt.Sprintf(SaveRet, l.VarStack[name]),
		)
	}
}

func (l *ListListener) ExitAssign(c *parser.AssignContext) {
	//	fmt.Println("ExitAssign")

	name := c.IDENT().GetText()
	l.addVarToStack(name)
	l.initializeList(name)
	l.fillList(name, c.ListType())
}

func (l *ListListener) EnterFuncDec(c *parser.FuncDecContext) {
	fmt.Println("EnterDeclaration")

	funk := c.IDENT().GetText()
	fmt.Printf("fn: %s\n", funk)

	l.initFunk(funk)
}

func (l *ListListener) ExitFuncDec(c *parser.FuncDecContext) {
	l.decFunkStack()
}

func (l *ListListener) printList(ident string) {
	if !l.inStack(ident) {
		log.Panicf("Unknown identifier '%s'\n", ident)
	} else {
		l.addToFunkBody(
			fmt.Sprintf(Addr2Arg1, l.VarStack[ident]) +
				PrintList,
		)
	}
}

func (l *ListListener) sendListOut(send string, recv string, direction rune) {
	if !l.inStack(send) {
		log.Panicf("Unknown identifier '%s'\n", send)
	} else {
		l.addToFunkBody(
			fmt.Sprintf(Addr2Arg1, l.VarStack[send]) +
				fmt.Sprintf(Int2Arg2, fileDescriptors[recv]),
		)

		switch direction {
		case '<':
			l.addToFunkBody(SendListL)
		case '>':
			l.addToFunkBody(SendListR)
		default:
			log.Panic("Unrecognized direction %c\n", direction)
		}
	}
}

func (l *ListListener) Send(send, recv string, direction rune) {
	if isFdMacro(recv) {
		switch fileDescriptors[recv] {
		case StdOut:
			l.sendListOut(send, recv, direction)
		default:
			fmt.Printf("%s not implemented\n", recv)
		}
	}
}

func (l *ListListener) ExitSendLeft(c *parser.SendLeftContext) {
	fmt.Println("ExitSendLeft")

	nId := len(c.AllIterable()) - 1
	for ; nId >= 0; nId-- {
		id := c.Iterable(nId)
		fmt.Printf("\titer: %s\n", id.GetText())
	}
	recv, send := c.Iterable(0).GetText(), c.Iterable(1).GetText()
	l.Send(send, recv, '<')
}

func (l *ListListener) ExitSendRight(c *parser.SendRightContext) {
	//	fmt.Println("ExitSendRight")

	send, recv := c.Iterable(0).GetText(), c.Iterable(1).GetText()
	l.Send(send, recv, '>')
}

func (l *ListListener) ExitStart(c *parser.StartContext) {
	l.Funks["main"].AsmPrelude = fmt.Sprintf(Prelude, l.StackSize)

	for _, a := range l.VarStack {
		l.Funks["main"].AsmEpilogue += fmt.Sprintf(Addr2Arg1, a)
		l.Funks["main"].AsmEpilogue += FreeLists
	}
	l.Funks["main"].AsmEpilogue += End
}

func (l *ListListener) GetAsm() (fullAsm string) {
	for _, asm := range l.Funks {
		fullAsm += asm.AsmPrelude
		fullAsm += asm.AsmBody
		fullAsm += asm.AsmEpilogue
	}
	return fullAsm
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
	//	fmt.Println("EnterNumber")
	i, err := strconv.Atoi(c.GetText())
	HandleErr(err)
	l.push(i)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (l *ListListener) ExitExprList(c *parser.ExprListContext) {
	/*	fmt.Println("ExitExprList")
		for _, e := range c.AllExpr() {
			fmt.Printf("e: %s\n", e.GetText())
		} */
}
