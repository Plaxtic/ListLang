// List.g4
grammar List;

// Tokens
LT:     '<';
GT:     '>';
MUL:    '*';
DIV:    '/';
ADD:    '+';
SUB:    '-';
BANG:   '!';
ASSIGN: '=';
COMMA:  ',';
QMARK:  '?';
COLON:  ':';
QUOTE:  '"';
MOD:    '%';

LSQUARE: '[';
RSQUARE: ']';
LPAREN:  '(';
RPAREN:  ')';
LBRACK:  '{';
RBRACK:  '}';

EQ:    '==';
N_EQ:  '!=';
LTE:   '<=';
GTE:   '>=';
LSEND: '<-';
RSEND: '->';

COMMENT: '//';

TRUE:  'true';
FALSE: 'false';

FUNC: 'fn';

MACRO: [A-Z][a-z]+;
IDENT: [a-z]+;
INT: [0-9]+;
WHITESPACE: [ \r\t] -> skip;

NL: '\n';

// Rules
start
    : line (NL line)* EOF
    ;

line
    : statExpr*
    | comment
    | NL
    ;

statExpr
    : state
    | expr
    ;

iterable
    : IDENT
    | MACRO
    | listType
    ;

state
    : iterable (LSEND iterable)+ # SendLeft
    | iterable (RSEND iterable)+ # SendRight
    | funcDec                    # Declaration
    ;


expr
    : expr op=(MUL|DIV|MOD) expr     # MulDivMod
    | expr op=(SUB|ADD) expr         # SubAdd
    | expr op=(EQ|N_EQ|LTE|GTE) expr # Cond
    | INT                            # Number 
    | IDENT ASSIGN listType          # Assign
    | IDENT                          # Variable

//    | listType                       # list
//    | LPAREN expr RPAREN             # Parenthesis
//    | expr CONDITIONAL expr          # Comp
//    | BANG expr                      # Not
//    | expr QMARK expr COLON expr     # Cond
    ;

ops: (MUL|DIV|SUB|ADD|MOD);
comment: '//' ().*? ;         // shite

args
    : exprList
    ;

funcDec
    : FUNC IDENT COLON ( args )? LBRACK (NL (statExpr|comment)*)* RBRACK
    ;

exprList
    : expr ( COMMA expr )*
    ;

listType
    : LSQUARE ( exprList )? RSQUARE
    ;
