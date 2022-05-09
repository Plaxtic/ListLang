package compiler

const (
	QWordSize = 0x8
	RetAndRpb = QWordSize * 2
	Prelude   = `
BITS 64

%%include "data_structures/list.asm"
%%include "data_structures/binarray.asm"

global _start

_start:
    mov rdi, [rsp]       ; argc
    mov rsi, [rsp + 0x8] ; argv
    call main

    mov rdi, rax
    mov rax, 0x3c        ; exit
    syscall

main:
    push rbp
    mov rbp, rsp
    sub rsp, %#x
    `
	FnPrelude = `
%s:
    push rbp
    mov rbp, rsp
    sub rsp, %#x
	`
	End = `
    xor rax, rax
    leave
    ret
`
	DivTwoVar = `
	mov rax, [rbp + %#x]
	mov r8, [rbp + %#x]
	div r8
	`
	DivOneVarLeft = `
	mov rax, %#x
	mov r8, [rbp + %#x]
	div r8
	`
	DivOneVarRight = `
	mov rax, [rbp + %#x]
	mov r8, %#x
	div r8
	`
	MulTwoVar = `
	mov rax, [rbp + %#x]
	imul rax, [rbp + %#x]
	`
	MulOneVarRight = `
	mov rax, %#x
	imul rax, [rbp + %#x]
	`
	MulOneVarLeft = `
	mov rax, [rbp + %#x]
	imul rax, %#x
	`
	AddTwoVar = `
	mov rax, [rbp + %#x]
	add rax, [rbp + %#x]
	`
	AddOneVarRight = `
	mov rax, %#x
	add rax, [rbp + %#x]
	`
	AddOneVarLeft = `
	mov rax, [rbp + %#x]
	add rax, %#x
	`
	SubTwoVar = `
	mov rax, [rbp + %#x]
	sub rax, [rbp + %#x]
	`
	SubOneVarRight = `
	mov rax, %#x
	sub rax, [rbp + %#x]
	`
	SubOneVarLeft = `
	mov rax, [rbp + %#x]
	sub rax, %#x
	`
	ListInit = `call list
	`
	Addr2Arg1 = `mov rdi, [rbp + %#x]
    `
	Addr2Arg2 = `mov rsi, [rbp + %#x]
    `
	Int2Arg2 = `mov rsi, %#x
    `
	Ret2Arg1 = `mov rdi, rax
	`
	Ret2Arg2 = `mov rsi, rax
	`
	SaveRet = `mov [rbp + %#x], rax
    `
	AppendList = `
    mov rsi, %#x
    call list_append
    mov rdi, rax
    `
	FreeLists = `call free_list
    `
	PrintList = `call list_print
    `
	SendListL = `call list_send_fd_L
    `
	SendListR = `call list_send_fd_R
    `
	JoinList = `call list_join
	`

	StdIn  = 0
	StdOut = 1
	StdErr = 2
)

var keyWords = [...]string{
	"print",
}

var fileDescriptors = map[string]int{
	"In":  0,
	"Out": 1,
	"Err": 2,
}

func isKeyWord(ident string) bool {
	for _, keyWord := range keyWords {
		if ident == keyWord {
			return true
		}
	}
	return false
}

func isFdMacro(ident string) bool {
	if _, ok := fileDescriptors[ident]; ok {
		return true
	}
	return false
}
