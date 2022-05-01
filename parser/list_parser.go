// Code generated from List.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // List

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ListParser struct {
	*antlr.BaseParser
}

var listParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func listParserInit() {
	staticData := &listParserStaticData
	staticData.literalNames = []string{
		"", "'<'", "'>'", "'*'", "'/'", "'+'", "'-'", "'!'", "'='", "','", "'?'",
		"':'", "'\"'", "'['", "']'", "'('", "')'", "'{'", "'}'", "'=='", "'!='",
		"'<='", "'>='", "'<-'", "'->'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA",
		"QMARK", "COLON", "QUOTE", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "TRUE",
		"FALSE", "IDENT", "INT", "WHITESPACE", "CONDITIONAL",
	}
	staticData.ruleNames = []string{
		"start", "expr", "exprList", "listType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 66, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 33, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 2, 1, 2, 1, 2,
		5, 2, 55, 8, 2, 10, 2, 12, 2, 58, 9, 2, 1, 3, 1, 3, 3, 3, 62, 8, 3, 1,
		3, 1, 3, 1, 3, 0, 1, 2, 4, 0, 2, 4, 6, 0, 1, 1, 0, 3, 6, 73, 0, 11, 1,
		0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 10,
		3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 11,
		12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 14, 15, 5, 0, 0,
		1, 15, 1, 1, 0, 0, 0, 16, 17, 6, 1, -1, 0, 17, 33, 5, 28, 0, 0, 18, 33,
		5, 27, 0, 0, 19, 33, 3, 6, 3, 0, 20, 21, 5, 15, 0, 0, 21, 22, 3, 2, 1,
		0, 22, 23, 5, 16, 0, 0, 23, 33, 1, 0, 0, 0, 24, 25, 5, 27, 0, 0, 25, 26,
		5, 8, 0, 0, 26, 33, 3, 6, 3, 0, 27, 28, 5, 7, 0, 0, 28, 33, 3, 2, 1, 3,
		29, 30, 5, 27, 0, 0, 30, 31, 5, 23, 0, 0, 31, 33, 3, 2, 1, 1, 32, 16, 1,
		0, 0, 0, 32, 18, 1, 0, 0, 0, 32, 19, 1, 0, 0, 0, 32, 20, 1, 0, 0, 0, 32,
		24, 1, 0, 0, 0, 32, 27, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 33, 48, 1, 0, 0,
		0, 34, 35, 10, 10, 0, 0, 35, 36, 7, 0, 0, 0, 36, 47, 3, 2, 1, 11, 37, 38,
		10, 4, 0, 0, 38, 39, 5, 30, 0, 0, 39, 47, 3, 2, 1, 5, 40, 41, 10, 2, 0,
		0, 41, 42, 5, 10, 0, 0, 42, 43, 3, 2, 1, 0, 43, 44, 5, 11, 0, 0, 44, 45,
		3, 2, 1, 3, 45, 47, 1, 0, 0, 0, 46, 34, 1, 0, 0, 0, 46, 37, 1, 0, 0, 0,
		46, 40, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1,
		0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 56, 3, 2, 1, 0, 52,
		53, 5, 9, 0, 0, 53, 55, 3, 2, 1, 0, 54, 52, 1, 0, 0, 0, 55, 58, 1, 0, 0,
		0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 5, 1, 0, 0, 0, 58, 56, 1,
		0, 0, 0, 59, 61, 5, 13, 0, 0, 60, 62, 3, 4, 2, 0, 61, 60, 1, 0, 0, 0, 61,
		62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5, 14, 0, 0, 64, 7, 1, 0, 0,
		0, 6, 11, 32, 46, 48, 56, 61,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ListParserInit initializes any static state used to implement ListParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewListParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ListParserInit() {
	staticData := &listParserStaticData
	staticData.once.Do(listParserInit)
}

// NewListParser produces a new parser instance for the optional input antlr.TokenStream.
func NewListParser(input antlr.TokenStream) *ListParser {
	ListParserInit()
	this := new(ListParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &listParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "List.g4"

	return this
}

// ListParser tokens.
const (
	ListParserEOF         = antlr.TokenEOF
	ListParserLT          = 1
	ListParserGT          = 2
	ListParserMUL         = 3
	ListParserDIV         = 4
	ListParserADD         = 5
	ListParserSUB         = 6
	ListParserBANG        = 7
	ListParserASSIGN      = 8
	ListParserCOMMA       = 9
	ListParserQMARK       = 10
	ListParserCOLON       = 11
	ListParserQUOTE       = 12
	ListParserLSQUARE     = 13
	ListParserRSQUARE     = 14
	ListParserLPAREN      = 15
	ListParserRPAREN      = 16
	ListParserLBRACK      = 17
	ListParserRBRACK      = 18
	ListParserEQ          = 19
	ListParserN_EQ        = 20
	ListParserLTE         = 21
	ListParserGTE         = 22
	ListParserLSEND       = 23
	ListParserRSEND       = 24
	ListParserTRUE        = 25
	ListParserFALSE       = 26
	ListParserIDENT       = 27
	ListParserINT         = 28
	ListParserWHITESPACE  = 29
	ListParserCONDITIONAL = 30
)

// ListParser rules.
const (
	ListParserRULE_start    = 0
	ListParserRULE_expr     = 1
	ListParserRULE_exprList = 2
	ListParserRULE_listType = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ListParserEOF, 0)
}

func (s *StartContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ListParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ListParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserBANG)|(1<<ListParserLSQUARE)|(1<<ListParserLPAREN)|(1<<ListParserIDENT)|(1<<ListParserINT))) != 0 {
		{
			p.SetState(8)
			p.expr(0)
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(ListParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesisContext struct {
	*ExprContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ListParserLPAREN, 0)
}

func (s *ParenthesisContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ListParserRPAREN, 0)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type NotContext struct {
	*ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) BANG() antlr.TerminalNode {
	return s.GetToken(ListParserBANG, 0)
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitNot(s)
	}
}

type CondContext struct {
	*ExprContext
}

func NewCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondContext {
	var p = new(CondContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondContext) QMARK() antlr.TerminalNode {
	return s.GetToken(ListParserQMARK, 0)
}

func (s *CondContext) COLON() antlr.TerminalNode {
	return s.GetToken(ListParserCOLON, 0)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitCond(s)
	}
}

type VariableContext struct {
	*ExprContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitVariable(s)
	}
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(ListParserINT, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitNumber(s)
	}
}

type AssignContext struct {
	*ExprContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ListParserASSIGN, 0)
}

func (s *AssignContext) ListType() IListTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitAssign(s)
	}
}

type CalcContext struct {
	*ExprContext
	op antlr.Token
}

func NewCalcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalcContext {
	var p = new(CalcContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CalcContext) GetOp() antlr.Token { return s.op }

func (s *CalcContext) SetOp(v antlr.Token) { s.op = v }

func (s *CalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CalcContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CalcContext) MUL() antlr.TerminalNode {
	return s.GetToken(ListParserMUL, 0)
}

func (s *CalcContext) DIV() antlr.TerminalNode {
	return s.GetToken(ListParserDIV, 0)
}

func (s *CalcContext) SUB() antlr.TerminalNode {
	return s.GetToken(ListParserSUB, 0)
}

func (s *CalcContext) ADD() antlr.TerminalNode {
	return s.GetToken(ListParserADD, 0)
}

func (s *CalcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterCalc(s)
	}
}

func (s *CalcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitCalc(s)
	}
}

type ListContext struct {
	*ExprContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ListType() IListTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitList(s)
	}
}

type SendContext struct {
	*ExprContext
}

func NewSendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendContext {
	var p = new(SendContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *SendContext) LSEND() antlr.TerminalNode {
	return s.GetToken(ListParserLSEND, 0)
}

func (s *SendContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterSend(s)
	}
}

func (s *SendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitSend(s)
	}
}

type CompContext struct {
	*ExprContext
}

func NewCompContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompContext {
	var p = new(CompContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompContext) CONDITIONAL() antlr.TerminalNode {
	return s.GetToken(ListParserCONDITIONAL, 0)
}

func (s *CompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterComp(s)
	}
}

func (s *CompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitComp(s)
	}
}

func (p *ListParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ListParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ListParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(17)
			p.Match(ListParserINT)
		}

	case 2:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(ListParserIDENT)
		}

	case 3:
		localctx = NewListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.ListType()
		}

	case 4:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(ListParserLPAREN)
		}
		{
			p.SetState(21)
			p.expr(0)
		}
		{
			p.SetState(22)
			p.Match(ListParserRPAREN)
		}

	case 5:
		localctx = NewAssignContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(25)
			p.Match(ListParserASSIGN)
		}
		{
			p.SetState(26)
			p.ListType()
		}

	case 6:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.Match(ListParserBANG)
		}
		{
			p.SetState(28)
			p.expr(3)
		}

	case 7:
		localctx = NewSendContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(30)
			p.Match(ListParserLSEND)
		}
		{
			p.SetState(31)
			p.expr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCalcContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(35)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CalcContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserMUL)|(1<<ListParserDIV)|(1<<ListParserADD)|(1<<ListParserSUB))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CalcContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expr(11)
				}

			case 2:
				localctx = NewCompContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(38)
					p.Match(ListParserCONDITIONAL)
				}
				{
					p.SetState(39)
					p.expr(5)
				}

			case 3:
				localctx = NewCondContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(41)
					p.Match(ListParserQMARK)
				}
				{
					p.SetState(42)
					p.expr(0)
				}
				{
					p.SetState(43)
					p.Match(ListParserCOLON)
				}
				{
					p.SetState(44)
					p.expr(3)
				}

			}

		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ListParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ListParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *ListParser) ExprList() (localctx IExprListContext) {
	this := p
	_ = this

	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ListParserRULE_exprList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.expr(0)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserCOMMA {
		{
			p.SetState(52)
			p.Match(ListParserCOMMA)
		}
		{
			p.SetState(53)
			p.expr(0)
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(ListParserLSQUARE, 0)
}

func (s *ListTypeContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(ListParserRSQUARE, 0)
}

func (s *ListTypeContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *ListParser) ListType() (localctx IListTypeContext) {
	this := p
	_ = this

	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ListParserRULE_listType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(ListParserLSQUARE)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserBANG)|(1<<ListParserLSQUARE)|(1<<ListParserLPAREN)|(1<<ListParserIDENT)|(1<<ListParserINT))) != 0 {
		{
			p.SetState(60)
			p.ExprList()
		}

	}
	{
		p.SetState(63)
		p.Match(ListParserRSQUARE)
	}

	return localctx
}

func (p *ListParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ListParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
