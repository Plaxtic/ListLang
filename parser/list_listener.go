// Code generated from List.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // List

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ListListener is a complete listener for a parse tree produced by ListParser.
type ListListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterSendLeft is called when entering the SendLeft production.
	EnterSendLeft(c *SendLeftContext)

	// EnterSendRight is called when entering the SendRight production.
	EnterSendRight(c *SendRightContext)

	// EnterDeclaration is called when entering the Declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterSubAdd is called when entering the SubAdd production.
	EnterSubAdd(c *SubAddContext)

	// EnterCond is called when entering the Cond production.
	EnterCond(c *CondContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAssign is called when entering the Assign production.
	EnterAssign(c *AssignContext)

	// EnterFuncDec is called when entering the funcDec production.
	EnterFuncDec(c *FuncDecContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitSendLeft is called when exiting the SendLeft production.
	ExitSendLeft(c *SendLeftContext)

	// ExitSendRight is called when exiting the SendRight production.
	ExitSendRight(c *SendRightContext)

	// ExitDeclaration is called when exiting the Declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitSubAdd is called when exiting the SubAdd production.
	ExitSubAdd(c *SubAddContext)

	// ExitCond is called when exiting the Cond production.
	ExitCond(c *CondContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAssign is called when exiting the Assign production.
	ExitAssign(c *AssignContext)

	// ExitFuncDec is called when exiting the funcDec production.
	ExitFuncDec(c *FuncDecContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)
}
