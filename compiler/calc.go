package compiler

import (
	"fmt"
	"log"
	"strconv"
	"test/parser"
)

func (l *ListListener) topOfIntStack() int { return l.IntStack[len(l.IntStack)-1] }
func (l *ListListener) decIntStack()       { l.IntStack = l.IntStack[:len(l.IntStack)-1] }
func (l *ListListener) pushInt(i int)      { l.IntStack = append(l.IntStack, i) }

func (l *ListListener) popInt() (ret int) {
	if len(l.IntStack) < 1 {
		log.Panic("Cannot popInt: Stack Empty")
	}
	ret = l.topOfIntStack()
	l.decIntStack()

	return ret
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
			l.getVar(left),
			l.popInt(),
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
			l.popInt(),
			l.getVar(right),
		),
	)
}

func (l *ListListener) constMulDiv(opType int) {
	right, left := l.popInt(), l.popInt()

	switch opType {
	case parser.ListParserMUL:
		l.pushInt(left * right)
	case parser.ListParserDIV:
		l.pushInt(left / right)
	case parser.ListParserADD:
		l.pushInt(left + right)
	case parser.ListParserSUB:
		l.pushInt(left - right)
	}
}

func (l *ListListener) EnterNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	HandleErr(err)
	l.pushInt(i)
}
