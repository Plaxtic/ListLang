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
WHITESPACE: [ \r\n\t] -> skip;

CONDITIONAL: (EQ|N_EQ|LTE|GTE);

// Rules
start : expr* EOF;

//stmt
//    : IDENT '<-' expr # Send
//    | expr '->' IDENT # RSend
//    ;

expr
    : expr op=(MUL|DIV|SUB|ADD) expr # Calc
    | INT                            # Number 
    | IDENT                          # Variable
    | listType                       # list
    | LPAREN expr RPAREN             # Parenthesis
    | IDENT ASSIGN listType          # Assign
    | expr CONDITIONAL expr          # Comp
    | BANG expr                      # Not
    | expr QMARK expr COLON expr     # Cond
    | IDENT '<-' expr # Send
    ;

exprList: expr ( COMMA expr )*;
listType: LSQUARE ( exprList )? RSQUARE;
