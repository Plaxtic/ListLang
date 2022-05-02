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

TRUE:  'true';
FALSE: 'false';

IDENT: [a-z]+;
INT: [0-9]+;
WHITESPACE: [ \r\t] -> skip;

NL: '\n';

// Rules
start
    : line (NL line)* EOF
    ;

line
    : state*
    | expr*
    | NL
    ;


state
    : IDENT LSEND IDENT # SendLeft
    | IDENT RSEND IDENT # SendRight
    | funcDec           # Declaration
    ;

expr
    : expr op=(MUL|DIV) expr         # MulDiv
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

funcDec
    : IDENT LPAREN ( exprList )? RPAREN
    ;

exprList
    : expr ( COMMA expr )*
    ;

listType
    : LSQUARE ( exprList )? RSQUARE
    ;
