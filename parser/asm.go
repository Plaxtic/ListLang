package parser

var keyWords = [...]string{
	"for",
	"print",
}

const (
	ErrBuf        = 20
	QWordSize     = 0x8
	PreludeMaxLen = 314

	Prelude = `BITS 64

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

	End = `
    xor rax, rax
    leave
    ret
`

	ListinitMaxLen = 68
	ListInit       = `
    call list
    mov [rbp + %#x], rax		; %s
    `

	Addr2Arg1 = `mov rdi, [rbp + %#x]
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
)
