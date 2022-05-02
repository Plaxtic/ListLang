// Code generated from List.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ListLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var listlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func listlexerLexerInit() {
	staticData := &listlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA", "QMARK",
		"COLON", "QUOTE", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN", "LBRACK",
		"RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "TRUE", "FALSE",
		"IDENT", "INT", "WHITESPACE", "NL",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 142, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 4, 26, 128,
		8, 26, 11, 26, 12, 26, 129, 1, 27, 4, 27, 133, 8, 27, 11, 27, 12, 27, 134,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 0, 0, 30, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 1, 0, 3, 1,
		0, 97, 122, 1, 0, 48, 57, 3, 0, 9, 9, 13, 13, 32, 32, 143, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 63, 1,
		0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71,
		1, 0, 0, 0, 13, 73, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0,
		19, 79, 1, 0, 0, 0, 21, 81, 1, 0, 0, 0, 23, 83, 1, 0, 0, 0, 25, 85, 1,
		0, 0, 0, 27, 87, 1, 0, 0, 0, 29, 89, 1, 0, 0, 0, 31, 91, 1, 0, 0, 0, 33,
		93, 1, 0, 0, 0, 35, 95, 1, 0, 0, 0, 37, 97, 1, 0, 0, 0, 39, 100, 1, 0,
		0, 0, 41, 103, 1, 0, 0, 0, 43, 106, 1, 0, 0, 0, 45, 109, 1, 0, 0, 0, 47,
		112, 1, 0, 0, 0, 49, 115, 1, 0, 0, 0, 51, 120, 1, 0, 0, 0, 53, 127, 1,
		0, 0, 0, 55, 132, 1, 0, 0, 0, 57, 136, 1, 0, 0, 0, 59, 140, 1, 0, 0, 0,
		61, 62, 5, 60, 0, 0, 62, 2, 1, 0, 0, 0, 63, 64, 5, 62, 0, 0, 64, 4, 1,
		0, 0, 0, 65, 66, 5, 42, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 5, 47, 0, 0, 68,
		8, 1, 0, 0, 0, 69, 70, 5, 43, 0, 0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 45,
		0, 0, 72, 12, 1, 0, 0, 0, 73, 74, 5, 33, 0, 0, 74, 14, 1, 0, 0, 0, 75,
		76, 5, 61, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 5, 44, 0, 0, 78, 18, 1, 0,
		0, 0, 79, 80, 5, 63, 0, 0, 80, 20, 1, 0, 0, 0, 81, 82, 5, 58, 0, 0, 82,
		22, 1, 0, 0, 0, 83, 84, 5, 34, 0, 0, 84, 24, 1, 0, 0, 0, 85, 86, 5, 91,
		0, 0, 86, 26, 1, 0, 0, 0, 87, 88, 5, 93, 0, 0, 88, 28, 1, 0, 0, 0, 89,
		90, 5, 40, 0, 0, 90, 30, 1, 0, 0, 0, 91, 92, 5, 41, 0, 0, 92, 32, 1, 0,
		0, 0, 93, 94, 5, 123, 0, 0, 94, 34, 1, 0, 0, 0, 95, 96, 5, 125, 0, 0, 96,
		36, 1, 0, 0, 0, 97, 98, 5, 61, 0, 0, 98, 99, 5, 61, 0, 0, 99, 38, 1, 0,
		0, 0, 100, 101, 5, 33, 0, 0, 101, 102, 5, 61, 0, 0, 102, 40, 1, 0, 0, 0,
		103, 104, 5, 60, 0, 0, 104, 105, 5, 61, 0, 0, 105, 42, 1, 0, 0, 0, 106,
		107, 5, 62, 0, 0, 107, 108, 5, 61, 0, 0, 108, 44, 1, 0, 0, 0, 109, 110,
		5, 60, 0, 0, 110, 111, 5, 45, 0, 0, 111, 46, 1, 0, 0, 0, 112, 113, 5, 45,
		0, 0, 113, 114, 5, 62, 0, 0, 114, 48, 1, 0, 0, 0, 115, 116, 5, 116, 0,
		0, 116, 117, 5, 114, 0, 0, 117, 118, 5, 117, 0, 0, 118, 119, 5, 101, 0,
		0, 119, 50, 1, 0, 0, 0, 120, 121, 5, 102, 0, 0, 121, 122, 5, 97, 0, 0,
		122, 123, 5, 108, 0, 0, 123, 124, 5, 115, 0, 0, 124, 125, 5, 101, 0, 0,
		125, 52, 1, 0, 0, 0, 126, 128, 7, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 54, 1, 0,
		0, 0, 131, 133, 7, 1, 0, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 56, 1, 0, 0, 0, 136, 137,
		7, 2, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 6, 28, 0, 0, 139, 58, 1, 0,
		0, 0, 140, 141, 5, 10, 0, 0, 141, 60, 1, 0, 0, 0, 3, 0, 129, 134, 1, 6,
		0, 0,
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

// ListLexerInit initializes any static state used to implement ListLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewListLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ListLexerInit() {
	staticData := &listlexerLexerStaticData
	staticData.once.Do(listlexerLexerInit)
}

// NewListLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewListLexer(input antlr.CharStream) *ListLexer {
	ListLexerInit()
	l := new(ListLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &listlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "List.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ListLexer tokens.
const (
	ListLexerLT         = 1
	ListLexerGT         = 2
	ListLexerMUL        = 3
	ListLexerDIV        = 4
	ListLexerADD        = 5
	ListLexerSUB        = 6
	ListLexerBANG       = 7
	ListLexerASSIGN     = 8
	ListLexerCOMMA      = 9
	ListLexerQMARK      = 10
	ListLexerCOLON      = 11
	ListLexerQUOTE      = 12
	ListLexerLSQUARE    = 13
	ListLexerRSQUARE    = 14
	ListLexerLPAREN     = 15
	ListLexerRPAREN     = 16
	ListLexerLBRACK     = 17
	ListLexerRBRACK     = 18
	ListLexerEQ         = 19
	ListLexerN_EQ       = 20
	ListLexerLTE        = 21
	ListLexerGTE        = 22
	ListLexerLSEND      = 23
	ListLexerRSEND      = 24
	ListLexerTRUE       = 25
	ListLexerFALSE      = 26
	ListLexerIDENT      = 27
	ListLexerINT        = 28
	ListLexerWHITESPACE = 29
	ListLexerNL         = 30
)
