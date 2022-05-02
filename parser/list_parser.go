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
		"'<='", "'>='", "'<-'", "'->'", "'true'", "'false'", "", "", "", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA",
		"QMARK", "COLON", "QUOTE", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "TRUE",
		"FALSE", "IDENT", "INT", "WHITESPACE", "NL",
	}
	staticData.ruleNames = []string{
		"start", "line", "state", "expr", "funcDec", "exprList", "listType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 92, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1,
		1, 1, 5, 1, 32, 8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 1, 3, 1, 38, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 55, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 66, 8, 3, 10, 3, 12, 3, 69, 9, 3, 1, 4, 1, 4, 1, 4,
		3, 4, 74, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 81, 8, 5, 10, 5, 12,
		5, 84, 9, 5, 1, 6, 1, 6, 3, 6, 88, 8, 6, 1, 6, 1, 6, 1, 6, 0, 1, 6, 7,
		0, 2, 4, 6, 8, 10, 12, 0, 3, 1, 0, 3, 4, 1, 0, 5, 6, 1, 0, 19, 22, 99,
		0, 14, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 54, 1, 0, 0,
		0, 8, 70, 1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 19, 3,
		2, 1, 0, 15, 16, 5, 30, 0, 0, 16, 18, 3, 2, 1, 0, 17, 15, 1, 0, 0, 0, 18,
		21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 22, 1, 0, 0,
		0, 21, 19, 1, 0, 0, 0, 22, 23, 5, 0, 0, 1, 23, 1, 1, 0, 0, 0, 24, 26, 3,
		4, 2, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27,
		28, 1, 0, 0, 0, 28, 38, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 32, 3, 6, 3,
		0, 31, 30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34,
		1, 0, 0, 0, 34, 38, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 38, 5, 30, 0, 0,
		37, 27, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 3, 1, 0,
		0, 0, 39, 40, 5, 27, 0, 0, 40, 41, 5, 23, 0, 0, 41, 47, 5, 27, 0, 0, 42,
		43, 5, 27, 0, 0, 43, 44, 5, 24, 0, 0, 44, 47, 5, 27, 0, 0, 45, 47, 3, 8,
		4, 0, 46, 39, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 5,
		1, 0, 0, 0, 48, 49, 6, 3, -1, 0, 49, 55, 5, 28, 0, 0, 50, 51, 5, 27, 0,
		0, 51, 52, 5, 8, 0, 0, 52, 55, 3, 12, 6, 0, 53, 55, 5, 27, 0, 0, 54, 48,
		1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 67, 1, 0, 0, 0,
		56, 57, 10, 6, 0, 0, 57, 58, 7, 0, 0, 0, 58, 66, 3, 6, 3, 7, 59, 60, 10,
		5, 0, 0, 60, 61, 7, 1, 0, 0, 61, 66, 3, 6, 3, 6, 62, 63, 10, 4, 0, 0, 63,
		64, 7, 2, 0, 0, 64, 66, 3, 6, 3, 5, 65, 56, 1, 0, 0, 0, 65, 59, 1, 0, 0,
		0, 65, 62, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68,
		1, 0, 0, 0, 68, 7, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5, 27, 0, 0,
		71, 73, 5, 15, 0, 0, 72, 74, 3, 10, 5, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 5, 16, 0, 0, 76, 9, 1, 0, 0, 0, 77,
		82, 3, 6, 3, 0, 78, 79, 5, 9, 0, 0, 79, 81, 3, 6, 3, 0, 80, 78, 1, 0, 0,
		0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 11,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 87, 5, 13, 0, 0, 86, 88, 3, 10, 5,
		0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90,
		5, 14, 0, 0, 90, 13, 1, 0, 0, 0, 11, 19, 27, 33, 37, 46, 54, 65, 67, 73,
		82, 87,
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
	ListParserEOF        = antlr.TokenEOF
	ListParserLT         = 1
	ListParserGT         = 2
	ListParserMUL        = 3
	ListParserDIV        = 4
	ListParserADD        = 5
	ListParserSUB        = 6
	ListParserBANG       = 7
	ListParserASSIGN     = 8
	ListParserCOMMA      = 9
	ListParserQMARK      = 10
	ListParserCOLON      = 11
	ListParserQUOTE      = 12
	ListParserLSQUARE    = 13
	ListParserRSQUARE    = 14
	ListParserLPAREN     = 15
	ListParserRPAREN     = 16
	ListParserLBRACK     = 17
	ListParserRBRACK     = 18
	ListParserEQ         = 19
	ListParserN_EQ       = 20
	ListParserLTE        = 21
	ListParserGTE        = 22
	ListParserLSEND      = 23
	ListParserRSEND      = 24
	ListParserTRUE       = 25
	ListParserFALSE      = 26
	ListParserIDENT      = 27
	ListParserINT        = 28
	ListParserWHITESPACE = 29
	ListParserNL         = 30
)

// ListParser rules.
const (
	ListParserRULE_start    = 0
	ListParserRULE_line     = 1
	ListParserRULE_state    = 2
	ListParserRULE_expr     = 3
	ListParserRULE_funcDec  = 4
	ListParserRULE_exprList = 5
	ListParserRULE_listType = 6
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

func (s *StartContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
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

	return t.(ILineContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ListParserEOF, 0)
}

func (s *StartContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ListParserNL)
}

func (s *StartContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ListParserNL, i)
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
	{
		p.SetState(14)
		p.Line()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserNL {
		{
			p.SetState(15)
			p.Match(ListParserNL)
		}
		{
			p.SetState(16)
			p.Line()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(ListParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) AllState() []IStateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateContext); ok {
			len++
		}
	}

	tst := make([]IStateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateContext); ok {
			tst[i] = t.(IStateContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) State(i int) IStateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateContext); ok {
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

	return t.(IStateContext)
}

func (s *LineContext) AllExpr() []IExprContext {
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

func (s *LineContext) Expr(i int) IExprContext {
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

func (s *LineContext) NL() antlr.TerminalNode {
	return s.GetToken(ListParserNL, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *ListParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ListParserRULE_line)
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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ListParserIDENT {
			{
				p.SetState(24)
				p.State()
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ListParserIDENT || _la == ListParserINT {
			{
				p.SetState(30)
				p.expr(0)
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Match(ListParserNL)
		}

	}

	return localctx
}

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateContext differentiates from other interfaces.
	IsStateContext()
}

type StateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateContext() *StateContext {
	var p = new(StateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_state
	return p
}

func (*StateContext) IsStateContext() {}

func NewStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateContext {
	var p = new(StateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_state

	return p
}

func (s *StateContext) GetParser() antlr.Parser { return s.parser }

func (s *StateContext) CopyFrom(ctx *StateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SendRightContext struct {
	*StateContext
}

func NewSendRightContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendRightContext {
	var p = new(SendRightContext)

	p.StateContext = NewEmptyStateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateContext))

	return p
}

func (s *SendRightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendRightContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(ListParserIDENT)
}

func (s *SendRightContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, i)
}

func (s *SendRightContext) RSEND() antlr.TerminalNode {
	return s.GetToken(ListParserRSEND, 0)
}

func (s *SendRightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterSendRight(s)
	}
}

func (s *SendRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitSendRight(s)
	}
}

type DeclarationContext struct {
	*StateContext
}

func NewDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationContext {
	var p = new(DeclarationContext)

	p.StateContext = NewEmptyStateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateContext))

	return p
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) FuncDec() IFuncDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDecContext)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

type SendLeftContext struct {
	*StateContext
}

func NewSendLeftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendLeftContext {
	var p = new(SendLeftContext)

	p.StateContext = NewEmptyStateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateContext))

	return p
}

func (s *SendLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendLeftContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(ListParserIDENT)
}

func (s *SendLeftContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, i)
}

func (s *SendLeftContext) LSEND() antlr.TerminalNode {
	return s.GetToken(ListParserLSEND, 0)
}

func (s *SendLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterSendLeft(s)
	}
}

func (s *SendLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitSendLeft(s)
	}
}

func (p *ListParser) State() (localctx IStateContext) {
	this := p
	_ = this

	localctx = NewStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ListParserRULE_state)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSendLeftContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(40)
			p.Match(ListParserLSEND)
		}
		{
			p.SetState(41)
			p.Match(ListParserIDENT)
		}

	case 2:
		localctx = NewSendRightContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(43)
			p.Match(ListParserRSEND)
		}
		{
			p.SetState(44)
			p.Match(ListParserIDENT)
		}

	case 3:
		localctx = NewDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.FuncDec()
		}

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

type SubAddContext struct {
	*ExprContext
	op antlr.Token
}

func NewSubAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubAddContext {
	var p = new(SubAddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubAddContext) GetOp() antlr.Token { return s.op }

func (s *SubAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *SubAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAddContext) AllExpr() []IExprContext {
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

func (s *SubAddContext) Expr(i int) IExprContext {
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

func (s *SubAddContext) SUB() antlr.TerminalNode {
	return s.GetToken(ListParserSUB, 0)
}

func (s *SubAddContext) ADD() antlr.TerminalNode {
	return s.GetToken(ListParserADD, 0)
}

func (s *SubAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterSubAdd(s)
	}
}

func (s *SubAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitSubAdd(s)
	}
}

type CondContext struct {
	*ExprContext
	op antlr.Token
}

func NewCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondContext {
	var p = new(CondContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CondContext) GetOp() antlr.Token { return s.op }

func (s *CondContext) SetOp(v antlr.Token) { s.op = v }

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

func (s *CondContext) EQ() antlr.TerminalNode {
	return s.GetToken(ListParserEQ, 0)
}

func (s *CondContext) N_EQ() antlr.TerminalNode {
	return s.GetToken(ListParserN_EQ, 0)
}

func (s *CondContext) LTE() antlr.TerminalNode {
	return s.GetToken(ListParserLTE, 0)
}

func (s *CondContext) GTE() antlr.TerminalNode {
	return s.GetToken(ListParserGTE, 0)
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

type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
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

func (s *MulDivContext) Expr(i int) IExprContext {
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

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(ListParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(ListParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitMulDiv(s)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, ListParserRULE_expr, _p)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(49)
			p.Match(ListParserINT)
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(51)
			p.Match(ListParserASSIGN)
		}
		{
			p.SetState(52)
			p.ListType()
		}

	case 3:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(ListParserIDENT)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(57)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ListParserMUL || _la == ListParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)
					p.expr(7)
				}

			case 2:
				localctx = NewSubAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(60)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SubAddContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ListParserADD || _la == ListParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SubAddContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expr(6)
				}

			case 3:
				localctx = NewCondContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(63)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CondContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserEQ)|(1<<ListParserN_EQ)|(1<<ListParserLTE)|(1<<ListParserGTE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CondContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(64)
					p.expr(5)
				}

			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncDecContext is an interface to support dynamic dispatch.
type IFuncDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDecContext differentiates from other interfaces.
	IsFuncDecContext()
}

type FuncDecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDecContext() *FuncDecContext {
	var p = new(FuncDecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_funcDec
	return p
}

func (*FuncDecContext) IsFuncDecContext() {}

func NewFuncDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDecContext {
	var p = new(FuncDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_funcDec

	return p
}

func (s *FuncDecContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *FuncDecContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ListParserLPAREN, 0)
}

func (s *FuncDecContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ListParserRPAREN, 0)
}

func (s *FuncDecContext) ExprList() IExprListContext {
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

func (s *FuncDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterFuncDec(s)
	}
}

func (s *FuncDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitFuncDec(s)
	}
}

func (p *ListParser) FuncDec() (localctx IFuncDecContext) {
	this := p
	_ = this

	localctx = NewFuncDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ListParserRULE_funcDec)
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
		p.SetState(70)
		p.Match(ListParserIDENT)
	}
	{
		p.SetState(71)
		p.Match(ListParserLPAREN)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ListParserIDENT || _la == ListParserINT {
		{
			p.SetState(72)
			p.ExprList()
		}

	}
	{
		p.SetState(75)
		p.Match(ListParserRPAREN)
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
	p.EnterRule(localctx, 10, ListParserRULE_exprList)
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
		p.SetState(77)
		p.expr(0)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserCOMMA {
		{
			p.SetState(78)
			p.Match(ListParserCOMMA)
		}
		{
			p.SetState(79)
			p.expr(0)
		}

		p.SetState(84)
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
	p.EnterRule(localctx, 12, ListParserRULE_listType)
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
		p.SetState(85)
		p.Match(ListParserLSQUARE)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ListParserIDENT || _la == ListParserINT {
		{
			p.SetState(86)
			p.ExprList()
		}

	}
	{
		p.SetState(89)
		p.Match(ListParserRSQUARE)
	}

	return localctx
}

func (p *ListParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
