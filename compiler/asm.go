package compiler

const (
	QWordSize = 0x8
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
	End = `
    xor rax, rax
    leave
    ret
`
	ListInit = `
    call list
    mov [rbp + %#x], rax		; %s
    `
	Addr2Arg1 = `mov rdi, [rbp + %#x]
    `
	Int2Arg2 = `mov rsi, %#x
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
)

var keyWords = [...]string{
	"print",
}

var fileDescriptors = map[string]int{
	"stdin":  0,
	"stdout": 1,
	"stderr": 2,
}

func isKeyWord(ident string) bool {
	for _, keyWord := range keyWords {
		if ident == keyWord {
			return true
		}
	}
	return false
}

func isConstFd(ident string) bool {
	if _, ok := fileDescriptors[ident]; ok {
		return true
	}
	return false
}
