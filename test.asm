
BITS 64

%include "data_structures/list.asm"
%include "data_structures/binarray.asm"

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
    sub rsp, 0x48
    
    call list
    mov [rbp + 0x10], rax		; one
    
    call list
    mov [rbp + 0x18], rax		; two
    mov rdi, [rbp + 0x18]
    
    mov rsi, 0x5
    call list_append
    mov rdi, rax
    mov [rbp + 0x18], rax
    
    call list
    mov [rbp + 0x20], rax		; three
    
    call list
    mov [rbp + 0x28], rax		; four
    
    call list
    mov [rbp + 0x30], rax		; five
    mov rdi, [rbp + 0x30]
    
    mov rsi, 0x1
    call list_append
    mov rdi, rax
    
    mov rsi, 0x2
    call list_append
    mov rdi, rax
    
    mov rsi, 0x3
    call list_append
    mov rdi, rax
    
    mov rsi, 0x4
    call list_append
    mov rdi, rax
    mov [rbp + 0x30], rax
    
    call list
    mov [rbp + 0x38], rax		; nine
    mov rdi, [rbp + 0x38]
    
    mov rsi, 0x2
    call list_append
    mov rdi, rax
    
    mov rsi, -0x5
    call list_append
    mov rdi, rax
    
    mov rsi, 0x4
    call list_append
    mov rdi, rax
    
    mov rsi, 0x5
    call list_append
    mov rdi, rax
    
    mov rsi, 0x6
    call list_append
    mov rdi, rax
    
    mov rsi, 0xea
    call list_append
    mov rdi, rax
    
    mov rsi, 0x4f2e4
    call list_append
    mov rdi, rax
    
    mov rsi, 0x1
    call list_append
    mov rdi, rax
    
    mov rsi, 0x2
    call list_append
    mov rdi, rax
    
    mov rsi, 0x3
    call list_append
    mov rdi, rax
    
    mov rsi, 0x4
    call list_append
    mov rdi, rax
    
    mov rsi, 0x5
    call list_append
    mov rdi, rax
    
    mov rsi, 0x6
    call list_append
    mov rdi, rax
    mov [rbp + 0x38], rax
    
    call list
    mov [rbp + 0x40], rax		; test
    mov rdi, [rbp + 0x40]
    
    mov rsi, 0x1
    call list_append
    mov rdi, rax
    
    mov rsi, 0x4
    call list_append
    mov rdi, rax
    mov [rbp + 0x40], rax
    mov rdi, [rbp + 0x30]
    mov rsi, 0x1
    call list_send_fd_R
    mov rdi, [rbp + 0x18]
    mov rsi, 0x1
    call list_send_fd_L
    mov rdi, [rbp + 0x30]
    mov rsi, 0x1
    call list_send_fd_L
    mov rdi, [rbp + 0x38]
    mov rsi, 0x1
    call list_send_fd_L
    mov rdi, [rbp + 0x40]
    mov rsi, 0x1
    call list_send_fd_L
    mov rdi, [rbp + 0x20]
    mov rsi, 0x1
    call list_send_fd_L
    mov rdi, [rbp + 0x38]
    call free_list
    mov rdi, [rbp + 0x40]
    call free_list
    mov rdi, [rbp + 0x10]
    call free_list
    mov rdi, [rbp + 0x18]
    call free_list
    mov rdi, [rbp + 0x20]
    call free_list
    mov rdi, [rbp + 0x28]
    call free_list
    mov rdi, [rbp + 0x30]
    call free_list
    
    xor rax, rax
    leave
    ret
