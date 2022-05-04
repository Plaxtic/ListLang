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
		"'<='", "'>='", "'<-'", "'->'", "'true'", "'false'", "'fn'", "", "",
		"", "", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA",
		"QMARK", "COLON", "QUOTE", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "TRUE",
		"FALSE", "FUNC", "MACRO", "IDENT", "INT", "WHITESPACE", "NL",
	}
	staticData.ruleNames = []string{
		"start", "line", "statExpr", "iterable", "state", "expr", "funcDec",
		"exprList", "listType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 121, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 30, 8, 1,
		10, 1, 12, 1, 33, 9, 1, 1, 1, 3, 1, 36, 8, 1, 1, 2, 1, 2, 3, 2, 40, 8,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 45, 8, 3, 1, 4, 1, 4, 1, 4, 4, 4, 50, 8, 4,
		11, 4, 12, 4, 51, 1, 4, 1, 4, 1, 4, 4, 4, 57, 8, 4, 11, 4, 12, 4, 58, 1,
		4, 3, 4, 62, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 70, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 81, 8, 5, 10,
		5, 12, 5, 84, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 90, 8, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 95, 8, 6, 10, 6, 12, 6, 98, 9, 6, 5, 6, 100, 8, 6, 10, 6, 12,
		6, 103, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 110, 8, 7, 10, 7, 12,
		7, 113, 9, 7, 1, 8, 1, 8, 3, 8, 117, 8, 8, 1, 8, 1, 8, 1, 8, 0, 1, 10,
		9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 3, 1, 0, 3, 4, 1, 0, 5, 6, 1, 0, 19,
		22, 131, 0, 18, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 44,
		1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0,
		14, 106, 1, 0, 0, 0, 16, 114, 1, 0, 0, 0, 18, 23, 3, 2, 1, 0, 19, 20, 5,
		32, 0, 0, 20, 22, 3, 2, 1, 0, 21, 19, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23,
		21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 1, 0, 0, 0, 25, 23, 1, 0, 0,
		0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 30, 3, 4, 2, 0, 29, 28, 1,
		0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32,
		36, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 36, 5, 32, 0, 0, 35, 31, 1, 0,
		0, 0, 35, 34, 1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 40, 3, 8, 4, 0, 38, 40,
		3, 10, 5, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 5, 1, 0, 0, 0,
		41, 45, 5, 29, 0, 0, 42, 45, 5, 28, 0, 0, 43, 45, 3, 16, 8, 0, 44, 41,
		1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 7, 1, 0, 0, 0,
		46, 49, 3, 6, 3, 0, 47, 48, 5, 23, 0, 0, 48, 50, 3, 6, 3, 0, 49, 47, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		62, 1, 0, 0, 0, 53, 56, 3, 6, 3, 0, 54, 55, 5, 24, 0, 0, 55, 57, 3, 6,
		3, 0, 56, 54, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 62, 3, 12, 6, 0, 61, 46, 1, 0, 0, 0,
		61, 53, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 6, 5,
		-1, 0, 64, 70, 5, 30, 0, 0, 65, 66, 5, 29, 0, 0, 66, 67, 5, 8, 0, 0, 67,
		70, 3, 16, 8, 0, 68, 70, 5, 29, 0, 0, 69, 63, 1, 0, 0, 0, 69, 65, 1, 0,
		0, 0, 69, 68, 1, 0, 0, 0, 70, 82, 1, 0, 0, 0, 71, 72, 10, 6, 0, 0, 72,
		73, 7, 0, 0, 0, 73, 81, 3, 10, 5, 7, 74, 75, 10, 5, 0, 0, 75, 76, 7, 1,
		0, 0, 76, 81, 3, 10, 5, 6, 77, 78, 10, 4, 0, 0, 78, 79, 7, 2, 0, 0, 79,
		81, 3, 10, 5, 5, 80, 71, 1, 0, 0, 0, 80, 74, 1, 0, 0, 0, 80, 77, 1, 0,
		0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 11,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 27, 0, 0, 86, 87, 5, 29, 0,
		0, 87, 89, 5, 11, 0, 0, 88, 90, 3, 14, 7, 0, 89, 88, 1, 0, 0, 0, 89, 90,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 101, 5, 17, 0, 0, 92, 96, 5, 32, 0,
		0, 93, 95, 3, 4, 2, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0,
		99, 92, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 18,
		0, 0, 105, 13, 1, 0, 0, 0, 106, 111, 3, 10, 5, 0, 107, 108, 5, 9, 0, 0,
		108, 110, 3, 10, 5, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 15, 1, 0, 0, 0, 113, 111, 1,
		0, 0, 0, 114, 116, 5, 13, 0, 0, 115, 117, 3, 14, 7, 0, 116, 115, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 5, 14, 0, 0,
		119, 17, 1, 0, 0, 0, 16, 23, 31, 35, 39, 44, 51, 58, 61, 69, 80, 82, 89,
		96, 101, 111, 116,
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
	ListParserFUNC       = 27
	ListParserMACRO      = 28
	ListParserIDENT      = 29
	ListParserINT        = 30
	ListParserWHITESPACE = 31
	ListParserNL         = 32
)

// ListParser rules.
const (
	ListParserRULE_start    = 0
	ListParserRULE_line     = 1
	ListParserRULE_statExpr = 2
	ListParserRULE_iterable = 3
	ListParserRULE_state    = 4
	ListParserRULE_expr     = 5
	ListParserRULE_funcDec  = 6
	ListParserRULE_exprList = 7
	ListParserRULE_listType = 8
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
		p.SetState(18)
		p.Line()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserNL {
		{
			p.SetState(19)
			p.Match(ListParserNL)
		}
		{
			p.SetState(20)
			p.Line()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
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

func (s *LineContext) AllStatExpr() []IStatExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatExprContext); ok {
			len++
		}
	}

	tst := make([]IStatExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatExprContext); ok {
			tst[i] = t.(IStatExprContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) StatExpr(i int) IStatExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatExprContext); ok {
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

	return t.(IStatExprContext)
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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserLSQUARE)|(1<<ListParserFUNC)|(1<<ListParserMACRO)|(1<<ListParserIDENT)|(1<<ListParserINT))) != 0 {
			{
				p.SetState(28)
				p.StatExpr()
			}

			p.SetState(33)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(ListParserNL)
		}

	}

	return localctx
}

// IStatExprContext is an interface to support dynamic dispatch.
type IStatExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatExprContext differentiates from other interfaces.
	IsStatExprContext()
}

type StatExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatExprContext() *StatExprContext {
	var p = new(StatExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_statExpr
	return p
}

func (*StatExprContext) IsStatExprContext() {}

func NewStatExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatExprContext {
	var p = new(StatExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_statExpr

	return p
}

func (s *StatExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StatExprContext) State() IStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateContext)
}

func (s *StatExprContext) Expr() IExprContext {
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

func (s *StatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterStatExpr(s)
	}
}

func (s *StatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitStatExpr(s)
	}
}

func (p *ListParser) StatExpr() (localctx IStatExprContext) {
	this := p
	_ = this

	localctx = NewStatExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ListParserRULE_statExpr)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.State()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.expr(0)
		}

	}

	return localctx
}

// IIterableContext is an interface to support dynamic dispatch.
type IIterableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterableContext differentiates from other interfaces.
	IsIterableContext()
}

type IterableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterableContext() *IterableContext {
	var p = new(IterableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ListParserRULE_iterable
	return p
}

func (*IterableContext) IsIterableContext() {}

func NewIterableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterableContext {
	var p = new(IterableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ListParserRULE_iterable

	return p
}

func (s *IterableContext) GetParser() antlr.Parser { return s.parser }

func (s *IterableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *IterableContext) MACRO() antlr.TerminalNode {
	return s.GetToken(ListParserMACRO, 0)
}

func (s *IterableContext) ListType() IListTypeContext {
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

func (s *IterableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.EnterIterable(s)
	}
}

func (s *IterableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ListListener); ok {
		listenerT.ExitIterable(s)
	}
}

func (p *ListParser) Iterable() (localctx IIterableContext) {
	this := p
	_ = this

	localctx = NewIterableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ListParserRULE_iterable)

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

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ListParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(ListParserIDENT)
		}

	case ListParserMACRO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(ListParserMACRO)
		}

	case ListParserLSQUARE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.ListType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *SendRightContext) AllIterable() []IIterableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIterableContext); ok {
			len++
		}
	}

	tst := make([]IIterableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIterableContext); ok {
			tst[i] = t.(IIterableContext)
			i++
		}
	}

	return tst
}

func (s *SendRightContext) Iterable(i int) IIterableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterableContext); ok {
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

	return t.(IIterableContext)
}

func (s *SendRightContext) AllRSEND() []antlr.TerminalNode {
	return s.GetTokens(ListParserRSEND)
}

func (s *SendRightContext) RSEND(i int) antlr.TerminalNode {
	return s.GetToken(ListParserRSEND, i)
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

func (s *SendLeftContext) AllIterable() []IIterableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIterableContext); ok {
			len++
		}
	}

	tst := make([]IIterableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIterableContext); ok {
			tst[i] = t.(IIterableContext)
			i++
		}
	}

	return tst
}

func (s *SendLeftContext) Iterable(i int) IIterableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterableContext); ok {
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

	return t.(IIterableContext)
}

func (s *SendLeftContext) AllLSEND() []antlr.TerminalNode {
	return s.GetTokens(ListParserLSEND)
}

func (s *SendLeftContext) LSEND(i int) antlr.TerminalNode {
	return s.GetToken(ListParserLSEND, i)
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
	p.EnterRule(localctx, 8, ListParserRULE_state)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSendLeftContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Iterable()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ListParserLSEND {
			{
				p.SetState(47)
				p.Match(ListParserLSEND)
			}
			{
				p.SetState(48)
				p.Iterable()
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewSendRightContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Iterable()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ListParserRSEND {
			{
				p.SetState(54)
				p.Match(ListParserRSEND)
			}
			{
				p.SetState(55)
				p.Iterable()
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		localctx = NewDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ListParserRULE_expr, _p)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(64)
			p.Match(ListParserINT)
		}

	case 2:
		localctx = NewAssignContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(ListParserIDENT)
		}
		{
			p.SetState(66)
			p.Match(ListParserASSIGN)
		}
		{
			p.SetState(67)
			p.ListType()
		}

	case 3:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(ListParserIDENT)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(72)

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
					p.SetState(73)
					p.expr(7)
				}

			case 2:
				localctx = NewSubAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(75)

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
					p.SetState(76)
					p.expr(6)
				}

			case 3:
				localctx = NewCondContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ListParserRULE_expr)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(78)

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
					p.SetState(79)
					p.expr(5)
				}

			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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

func (s *FuncDecContext) FUNC() antlr.TerminalNode {
	return s.GetToken(ListParserFUNC, 0)
}

func (s *FuncDecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(ListParserIDENT, 0)
}

func (s *FuncDecContext) COLON() antlr.TerminalNode {
	return s.GetToken(ListParserCOLON, 0)
}

func (s *FuncDecContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ListParserLBRACK, 0)
}

func (s *FuncDecContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ListParserRBRACK, 0)
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

func (s *FuncDecContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ListParserNL)
}

func (s *FuncDecContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ListParserNL, i)
}

func (s *FuncDecContext) AllStatExpr() []IStatExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatExprContext); ok {
			len++
		}
	}

	tst := make([]IStatExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatExprContext); ok {
			tst[i] = t.(IStatExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncDecContext) StatExpr(i int) IStatExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatExprContext); ok {
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

	return t.(IStatExprContext)
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
	p.EnterRule(localctx, 12, ListParserRULE_funcDec)
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
		p.Match(ListParserFUNC)
	}
	{
		p.SetState(86)
		p.Match(ListParserIDENT)
	}
	{
		p.SetState(87)
		p.Match(ListParserCOLON)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ListParserIDENT || _la == ListParserINT {
		{
			p.SetState(88)
			p.ExprList()
		}

	}
	{
		p.SetState(91)
		p.Match(ListParserLBRACK)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserNL {
		{
			p.SetState(92)
			p.Match(ListParserNL)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ListParserLSQUARE)|(1<<ListParserFUNC)|(1<<ListParserMACRO)|(1<<ListParserIDENT)|(1<<ListParserINT))) != 0 {
			{
				p.SetState(93)
				p.StatExpr()
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(ListParserRBRACK)
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
	p.EnterRule(localctx, 14, ListParserRULE_exprList)
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
		p.SetState(106)
		p.expr(0)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ListParserCOMMA {
		{
			p.SetState(107)
			p.Match(ListParserCOMMA)
		}
		{
			p.SetState(108)
			p.expr(0)
		}

		p.SetState(113)
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
	p.EnterRule(localctx, 16, ListParserRULE_listType)
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
		p.SetState(114)
		p.Match(ListParserLSQUARE)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ListParserIDENT || _la == ListParserINT {
		{
			p.SetState(115)
			p.ExprList()
		}

	}
	{
		p.SetState(118)
		p.Match(ListParserRSQUARE)
	}

	return localctx
}

func (p *ListParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
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
