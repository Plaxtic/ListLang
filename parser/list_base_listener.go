// Code generated from List.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // List

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseListListener is a complete listener for a parse tree produced by ListParser.
type BaseListListener struct{}

var _ ListListener = &BaseListListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseListListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseListListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseListListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseListListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseListListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseListListener) ExitStart(ctx *StartContext) {}

// EnterLine is called when production line is entered.
func (s *BaseListListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseListListener) ExitLine(ctx *LineContext) {}

// EnterSendLeft is called when production SendLeft is entered.
func (s *BaseListListener) EnterSendLeft(ctx *SendLeftContext) {}

// ExitSendLeft is called when production SendLeft is exited.
func (s *BaseListListener) ExitSendLeft(ctx *SendLeftContext) {}

// EnterSendRight is called when production SendRight is entered.
func (s *BaseListListener) EnterSendRight(ctx *SendRightContext) {}

// ExitSendRight is called when production SendRight is exited.
func (s *BaseListListener) ExitSendRight(ctx *SendRightContext) {}

// EnterDeclaration is called when production Declaration is entered.
func (s *BaseListListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production Declaration is exited.
func (s *BaseListListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterSubAdd is called when production SubAdd is entered.
func (s *BaseListListener) EnterSubAdd(ctx *SubAddContext) {}

// ExitSubAdd is called when production SubAdd is exited.
func (s *BaseListListener) ExitSubAdd(ctx *SubAddContext) {}

// EnterCond is called when production Cond is entered.
func (s *BaseListListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production Cond is exited.
func (s *BaseListListener) ExitCond(ctx *CondContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseListListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseListListener) ExitVariable(ctx *VariableContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseListListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseListListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseListListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseListListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAssign is called when production Assign is entered.
func (s *BaseListListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production Assign is exited.
func (s *BaseListListener) ExitAssign(ctx *AssignContext) {}

// EnterFuncDec is called when production funcDec is entered.
func (s *BaseListListener) EnterFuncDec(ctx *FuncDecContext) {}

// ExitFuncDec is called when production funcDec is exited.
func (s *BaseListListener) ExitFuncDec(ctx *FuncDecContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseListListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseListListener) ExitExprList(ctx *ExprListContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseListListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseListListener) ExitListType(ctx *ListTypeContext) {}
