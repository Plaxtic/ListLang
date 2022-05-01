// Code generated from List.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // List

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ListListener is a complete listener for a parse tree produced by ListParser.
type ListListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterCond is called when entering the Cond production.
	EnterCond(c *CondContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAssign is called when entering the Assign production.
	EnterAssign(c *AssignContext)

	// EnterCalc is called when entering the Calc production.
	EnterCalc(c *CalcContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterSend is called when entering the Send production.
	EnterSend(c *SendContext)

	// EnterComp is called when entering the Comp production.
	EnterComp(c *CompContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitCond is called when exiting the Cond production.
	ExitCond(c *CondContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAssign is called when exiting the Assign production.
	ExitAssign(c *AssignContext)

	// ExitCalc is called when exiting the Calc production.
	ExitCalc(c *CalcContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitSend is called when exiting the Send production.
	ExitSend(c *SendContext)

	// ExitComp is called when exiting the Comp production.
	ExitComp(c *CompContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)
}
