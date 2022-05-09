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

// EnterStatExpr is called when production statExpr is entered.
func (s *BaseListListener) EnterStatExpr(ctx *StatExprContext) {}

// ExitStatExpr is called when production statExpr is exited.
func (s *BaseListListener) ExitStatExpr(ctx *StatExprContext) {}

// EnterIterable is called when production iterable is entered.
func (s *BaseListListener) EnterIterable(ctx *IterableContext) {}

// ExitIterable is called when production iterable is exited.
func (s *BaseListListener) ExitIterable(ctx *IterableContext) {}

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

// EnterReturn is called when production Return is entered.
func (s *BaseListListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production Return is exited.
func (s *BaseListListener) ExitReturn(ctx *ReturnContext) {}

// EnterSubAdd is called when production SubAdd is entered.
func (s *BaseListListener) EnterSubAdd(ctx *SubAddContext) {}

// ExitSubAdd is called when production SubAdd is exited.
func (s *BaseListListener) ExitSubAdd(ctx *SubAddContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseListListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseListListener) ExitMulDivMod(ctx *MulDivModContext) {}

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

// EnterAssign is called when production Assign is entered.
func (s *BaseListListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production Assign is exited.
func (s *BaseListListener) ExitAssign(ctx *AssignContext) {}

// EnterOps is called when production ops is entered.
func (s *BaseListListener) EnterOps(ctx *OpsContext) {}

// ExitOps is called when production ops is exited.
func (s *BaseListListener) ExitOps(ctx *OpsContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseListListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseListListener) ExitComment(ctx *CommentContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseListListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseListListener) ExitArgs(ctx *ArgsContext) {}

// EnterFuncDec is called when production funcDec is entered.
func (s *BaseListListener) EnterFuncDec(ctx *FuncDecContext) {}

// ExitFuncDec is called when production funcDec is exited.
func (s *BaseListListener) ExitFuncDec(ctx *FuncDecContext) {}

// EnterRet is called when production ret is entered.
func (s *BaseListListener) EnterRet(ctx *RetContext) {}

// ExitRet is called when production ret is exited.
func (s *BaseListListener) ExitRet(ctx *RetContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseListListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseListListener) ExitExprList(ctx *ExprListContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseListListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseListListener) ExitListType(ctx *ListTypeContext) {}
