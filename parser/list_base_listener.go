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

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseListListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseListListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseListListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseListListener) ExitNot(ctx *NotContext) {}

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

// EnterCalc is called when production Calc is entered.
func (s *BaseListListener) EnterCalc(ctx *CalcContext) {}

// ExitCalc is called when production Calc is exited.
func (s *BaseListListener) ExitCalc(ctx *CalcContext) {}

// EnterList is called when production list is entered.
func (s *BaseListListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseListListener) ExitList(ctx *ListContext) {}

// EnterSend is called when production Send is entered.
func (s *BaseListListener) EnterSend(ctx *SendContext) {}

// ExitSend is called when production Send is exited.
func (s *BaseListListener) ExitSend(ctx *SendContext) {}

// EnterComp is called when production Comp is entered.
func (s *BaseListListener) EnterComp(ctx *CompContext) {}

// ExitComp is called when production Comp is exited.
func (s *BaseListListener) ExitComp(ctx *CompContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseListListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseListListener) ExitExprList(ctx *ExprListContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseListListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseListListener) ExitListType(ctx *ListTypeContext) {}
