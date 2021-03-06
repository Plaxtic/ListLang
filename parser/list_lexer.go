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
		"", "'ret'", "'<'", "'>'", "'*'", "'/'", "'+'", "'-'", "'!'", "'='",
		"','", "'?'", "':'", "'\"'", "'%'", "'['", "']'", "'('", "')'", "'{'",
		"'}'", "'=='", "'!='", "'<='", "'>='", "'<-'", "'->'", "'//'", "'true'",
		"'false'", "'fn'", "", "", "", "", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "", "LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA",
		"QMARK", "COLON", "QUOTE", "MOD", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "COMMENT",
		"TRUE", "FALSE", "FUNC", "MACRO", "IDENT", "INT", "WHITESPACE", "NL",
	}
	staticData.ruleNames = []string{
		"T__0", "LT", "GT", "MUL", "DIV", "ADD", "SUB", "BANG", "ASSIGN", "COMMA",
		"QMARK", "COLON", "QUOTE", "MOD", "LSQUARE", "RSQUARE", "LPAREN", "RPAREN",
		"LBRACK", "RBRACK", "EQ", "N_EQ", "LTE", "GTE", "LSEND", "RSEND", "COMMENT",
		"TRUE", "FALSE", "FUNC", "MACRO", "IDENT", "INT", "WHITESPACE", "NL",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 170, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		4, 30, 151, 8, 30, 11, 30, 12, 30, 152, 1, 31, 4, 31, 156, 8, 31, 11, 31,
		12, 31, 157, 1, 32, 4, 32, 161, 8, 32, 11, 32, 12, 32, 162, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 0, 0, 35, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 1, 0, 4, 1, 0, 65, 90, 1, 0, 97, 122, 1, 0, 48, 57, 3,
		0, 9, 9, 13, 13, 32, 32, 172, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0,
		5, 77, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0, 11, 83, 1, 0,
		0, 0, 13, 85, 1, 0, 0, 0, 15, 87, 1, 0, 0, 0, 17, 89, 1, 0, 0, 0, 19, 91,
		1, 0, 0, 0, 21, 93, 1, 0, 0, 0, 23, 95, 1, 0, 0, 0, 25, 97, 1, 0, 0, 0,
		27, 99, 1, 0, 0, 0, 29, 101, 1, 0, 0, 0, 31, 103, 1, 0, 0, 0, 33, 105,
		1, 0, 0, 0, 35, 107, 1, 0, 0, 0, 37, 109, 1, 0, 0, 0, 39, 111, 1, 0, 0,
		0, 41, 113, 1, 0, 0, 0, 43, 116, 1, 0, 0, 0, 45, 119, 1, 0, 0, 0, 47, 122,
		1, 0, 0, 0, 49, 125, 1, 0, 0, 0, 51, 128, 1, 0, 0, 0, 53, 131, 1, 0, 0,
		0, 55, 134, 1, 0, 0, 0, 57, 139, 1, 0, 0, 0, 59, 145, 1, 0, 0, 0, 61, 148,
		1, 0, 0, 0, 63, 155, 1, 0, 0, 0, 65, 160, 1, 0, 0, 0, 67, 164, 1, 0, 0,
		0, 69, 168, 1, 0, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 101, 0, 0, 73,
		74, 5, 116, 0, 0, 74, 2, 1, 0, 0, 0, 75, 76, 5, 60, 0, 0, 76, 4, 1, 0,
		0, 0, 77, 78, 5, 62, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 42, 0, 0, 80,
		8, 1, 0, 0, 0, 81, 82, 5, 47, 0, 0, 82, 10, 1, 0, 0, 0, 83, 84, 5, 43,
		0, 0, 84, 12, 1, 0, 0, 0, 85, 86, 5, 45, 0, 0, 86, 14, 1, 0, 0, 0, 87,
		88, 5, 33, 0, 0, 88, 16, 1, 0, 0, 0, 89, 90, 5, 61, 0, 0, 90, 18, 1, 0,
		0, 0, 91, 92, 5, 44, 0, 0, 92, 20, 1, 0, 0, 0, 93, 94, 5, 63, 0, 0, 94,
		22, 1, 0, 0, 0, 95, 96, 5, 58, 0, 0, 96, 24, 1, 0, 0, 0, 97, 98, 5, 34,
		0, 0, 98, 26, 1, 0, 0, 0, 99, 100, 5, 37, 0, 0, 100, 28, 1, 0, 0, 0, 101,
		102, 5, 91, 0, 0, 102, 30, 1, 0, 0, 0, 103, 104, 5, 93, 0, 0, 104, 32,
		1, 0, 0, 0, 105, 106, 5, 40, 0, 0, 106, 34, 1, 0, 0, 0, 107, 108, 5, 41,
		0, 0, 108, 36, 1, 0, 0, 0, 109, 110, 5, 123, 0, 0, 110, 38, 1, 0, 0, 0,
		111, 112, 5, 125, 0, 0, 112, 40, 1, 0, 0, 0, 113, 114, 5, 61, 0, 0, 114,
		115, 5, 61, 0, 0, 115, 42, 1, 0, 0, 0, 116, 117, 5, 33, 0, 0, 117, 118,
		5, 61, 0, 0, 118, 44, 1, 0, 0, 0, 119, 120, 5, 60, 0, 0, 120, 121, 5, 61,
		0, 0, 121, 46, 1, 0, 0, 0, 122, 123, 5, 62, 0, 0, 123, 124, 5, 61, 0, 0,
		124, 48, 1, 0, 0, 0, 125, 126, 5, 60, 0, 0, 126, 127, 5, 45, 0, 0, 127,
		50, 1, 0, 0, 0, 128, 129, 5, 45, 0, 0, 129, 130, 5, 62, 0, 0, 130, 52,
		1, 0, 0, 0, 131, 132, 5, 47, 0, 0, 132, 133, 5, 47, 0, 0, 133, 54, 1, 0,
		0, 0, 134, 135, 5, 116, 0, 0, 135, 136, 5, 114, 0, 0, 136, 137, 5, 117,
		0, 0, 137, 138, 5, 101, 0, 0, 138, 56, 1, 0, 0, 0, 139, 140, 5, 102, 0,
		0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 115, 0,
		0, 143, 144, 5, 101, 0, 0, 144, 58, 1, 0, 0, 0, 145, 146, 5, 102, 0, 0,
		146, 147, 5, 110, 0, 0, 147, 60, 1, 0, 0, 0, 148, 150, 7, 0, 0, 0, 149,
		151, 7, 1, 0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 62, 1, 0, 0, 0, 154, 156, 7, 1,
		0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 64, 1, 0, 0, 0, 159, 161, 7, 2, 0, 0, 160, 159,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 66, 1, 0, 0, 0, 164, 165, 7, 3, 0, 0, 165, 166, 1, 0, 0, 0,
		166, 167, 6, 33, 0, 0, 167, 68, 1, 0, 0, 0, 168, 169, 5, 10, 0, 0, 169,
		70, 1, 0, 0, 0, 4, 0, 152, 157, 162, 1, 6, 0, 0,
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
	ListLexerT__0       = 1
	ListLexerLT         = 2
	ListLexerGT         = 3
	ListLexerMUL        = 4
	ListLexerDIV        = 5
	ListLexerADD        = 6
	ListLexerSUB        = 7
	ListLexerBANG       = 8
	ListLexerASSIGN     = 9
	ListLexerCOMMA      = 10
	ListLexerQMARK      = 11
	ListLexerCOLON      = 12
	ListLexerQUOTE      = 13
	ListLexerMOD        = 14
	ListLexerLSQUARE    = 15
	ListLexerRSQUARE    = 16
	ListLexerLPAREN     = 17
	ListLexerRPAREN     = 18
	ListLexerLBRACK     = 19
	ListLexerRBRACK     = 20
	ListLexerEQ         = 21
	ListLexerN_EQ       = 22
	ListLexerLTE        = 23
	ListLexerGTE        = 24
	ListLexerLSEND      = 25
	ListLexerRSEND      = 26
	ListLexerCOMMENT    = 27
	ListLexerTRUE       = 28
	ListLexerFALSE      = 29
	ListLexerFUNC       = 30
	ListLexerMACRO      = 31
	ListLexerIDENT      = 32
	ListLexerINT        = 33
	ListLexerWHITESPACE = 34
	ListLexerNL         = 35
)
