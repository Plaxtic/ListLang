package compiler

import (
	"fmt"
	"log"
	"test/parser"
)

func (l *ListListener) sendRawList2Fd(send, recv string, direction rune) {
	l.fillRawList(send)
	l.addToFunkBody(fmt.Sprintf(Int2Arg2, fileDescriptors[recv]))

	switch direction {
	case '<':
		l.addToFunkBody(SendListL)
	case '>':
		l.addToFunkBody(SendListR)
	default:
		log.Panic("Unrecognized direction %c\n", direction)
	}
	l.addToFunkBody(FreeLists)
}

func (l *ListListener) sendVarList2Fd(send, recv string, direction rune) {
	l.addToFunkBody(fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[send]))
	l.addToFunkBody(fmt.Sprintf(Int2Arg2, fileDescriptors[recv]))

	switch direction {
	case '<':
		l.addToFunkBody(SendListL)
	case '>':
		l.addToFunkBody(SendListR)
	default:
		log.Panic("Unrecognized direction %c\n", direction)
	}
}

func (l *ListListener) sendRawList2List(send, recv string, direction rune) {
	l.fillRawList(send)
	l.addToFunkBody(
		fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[recv]) +
			Ret2Arg2 + JoinList +
			fmt.Sprintf(SaveRet, l.currFunk().VarStack[recv]),
	)
}

func (l *ListListener) sendVarList2List(send, recv string, direction rune) {
	l.addToFunkBody(
		fmt.Sprintf(Addr2Arg1, l.currFunk().VarStack[recv]) +
			fmt.Sprintf(Addr2Arg2, l.currFunk().VarStack[send]) +
			JoinList +
			fmt.Sprintf(SaveRet, l.currFunk().VarStack[recv]),
	)
}

func (l *ListListener) sendList2List(send, recv string, direction rune) {
	switch true {
	case l.inStack(send):
		l.sendVarList2List(send, recv, direction)
	case isList(send):
		l.sendRawList2List(send, recv, direction)
	default:
		log.Panicf("undefined name %s", send)
	}
}

func (l *ListListener) sendFdMacro(send, recv string, direction rune) {
	switch true {
	case l.inStack(send):
		l.sendVarList2Fd(send, recv, direction)
	case isList(send):
		l.sendRawList2Fd(send, recv, direction)
	default:
		log.Panicf("undefined name %s", send)
	}
}

func (l *ListListener) Send(send, recv string, direction rune) {
	switch true {
	case isFdMacro(recv):
		l.sendFdMacro(send, recv, direction)
	case l.inStack(recv):
		l.sendList2List(send, recv, direction)
	}
}

func (l *ListListener) ExitSendLeft(c *parser.SendLeftContext) {

	/*
		nId := len(c.AllIterable()) - 1
		for ; nId >= 0; nId-- {
			id := c.Iterable(nId)
			fmt.Printf("\titer: %s\n", id.GetText())
		}
	*/

	recv, send := c.Iterable(0).GetText(), c.Iterable(1).GetText()
	l.Send(send, recv, '<')
}

func (l *ListListener) ExitSendRight(c *parser.SendRightContext) {
	send, recv := c.Iterable(0).GetText(), c.Iterable(1).GetText()
	l.Send(send, recv, '>')
}
