BITS 64


;--------- CONSTANTS ---------; 

; file syscall numbers
SYS_WRITE   equ 0x1         
SYS_READ    equ 0x0         
SYS_CLOSE   equ 0x3         
SYS_EXIT    equ 0x3c        
SYS_OPEN    equ 0x2         

; open macros
O_RDONLY equ 0x0
O_WRONLY equ 0x1
O_CREAT  equ 0x40
BUFSIZ   equ 0x2000

; list macros
INTSIZ   equ 8
CHUNKSIZ equ 256
LISTSIZ  equ CHUNKSIZ*INTSIZ
;--------- CONSTANTS ---------; 

%include "data_structures/stdlib.asm"

extern malloc
extern realloc
extern free

; arg1 = rdi
; arg2 = rsi
; arg3 = rdx
; arg4 = rcx
; arg5 = r8

;; struct list {
;;     int len
;;     int *items
;; }

;; struct list *init_list()
list:
    push rdi

    mov rdi, LISTSIZ+1  ; +1 for len
    call malloc
    test rax, rax
    jz list_fail

    mov qword [rax], 0   ; len 

list_fail:
    pop rdi
    ret


;; struct list *append(struct list *l, int val)
list_append:
    push rdi
    push r8

    ;; if l->len != 0 && (l->len)%CHUNKSIZ == 0: realloc()
    mov r8, [rdi]
    test r8, r8
    jz list_append_insert
    and r8, CHUNKSIZ-1      ; fast modulo
    test r8, r8
    jnz list_append_insert

    ;; increase space

    ;; rax = realloc(l, l->len+1+LISTSIZ)
    push rsi
    push rdi

    mov rsi, [rdi]
    inc rsi
    imul rsi, INTSIZ
    add rsi, LISTSIZ
    call realloc

    pop rdi
    pop rsi

    ;; if rax == NULL: ret NULL
    test rax, rax
    jz list_append_ret

    ;; else: l = rax
    mov rdi, rax

list_append_insert:
    mov r8, [rdi]
    inc r8 
    mov [rdi + r8*INTSIZ], rsi
    mov [rdi], r8
    mov rax, rdi
    
list_append_ret:
    pop r8
    pop rdi
    ret


;; int index(struct list *l, int val)
list_index:
    push r8
    push r9

    mov r8, [rdi]
    xor r9, r9
    inc r9

list_index_search:
    cmp [rdi + r9*INTSIZ], rsi
    je list_index_found
    inc r9
    cmp r9, r8
    jl list_index_search

    mov rax, -1
    jmp list_index_ret

list_index_found:
    dec r9
    mov rax, r9

list_index_ret:
    pop r9
    pop r8
    ret
    

;; int list_pop(struct list *l, int idx, int *err)
list_pop:
    push rdi
    push rsi

    ;; if idx > l->len: *err = -1; ret 
    cmp rsi, [rdi]
    jg list_pop_fail

    ;; i = l->len - idx - 1
    mov rcx, [rdi]
    sub rcx, rsi
    dec rcx

    ;; rdi = &l->items[idx]
    lea rdi, [rdi + 1 + rsi*INTSIZ]
    mov rax, [rdi]

    ;; while (i > 0)
list_shift:
    ;;      *rdi = *rdi + 1
    mov rsi, [rdi + INTSIZ]
    mov [rdi], rsi

    ;;      rdi++
    ;;      i--
    add rdi, INTSIZ
    loop list_shift

    mov qword [rdx], 0
    jmp list_pop_ret

list_pop_fail:
    mov qword [rdx], -1

list_pop_ret:
    pop rsi
    pop rdi
    ret


;; struct list *list_insert(struct list *l, int idx, int value)
list_insert:
    push rdi
    push rsi
    push r8
    push r9
    push r10

    mov rax, rdi

    mov r8, [rdi]
    cmp rsi, r8
    jl list_insert_within_bounds
    mov rsi, r8

    ;; if l->len != 0 && (l->len)%CHUNKSIZ == 0: realloc()
list_insert_within_bounds:
    test r8, r8
    jz list_insert_insert

    inc r8
    and r8, CHUNKSIZ-1      ; fast modulo
    test r8, r8
    jnz list_insert_insert

    ;; increase space
    ;; rax = realloc(l, l->len+1+LISTSIZ)
    push rsi

    mov rsi, [rdi]
    inc rsi
    imul rsi, INTSIZ
    add rsi, LISTSIZ
    call realloc
    test rax, rax
    jz list_insert_ret
    mov rdi, rax

    pop rsi

list_insert_insert:
    mov r8, [rdi]
    lea r8, [rdi + INTSIZ + r8*INTSIZ]   ; r8 = &l->items[-1]
    lea r9, [rdi + INTSIZ + rsi*INTSIZ]  ; r9 = &l->items[idx]

list_insert_shift: 
    mov r10, [r8]
    mov [r8 + INTSIZ], r10

    sub r8, INTSIZ
    cmp r8, r9
    jg list_insert_shift

    mov [r8], rdx

list_insert_ret:
    pop r10
    pop r9
    pop r8
    pop rsi
    pop rdi
    ret


;; void list_reverse(struct list *l)
list_reverse:
    push r8
    push r9
    push r10

    mov r8, [rdi]     ; len
    mov r9, r8
    imul r9, INTSIZ
    add r9, rdi       ; &l->items[-1]
    sar r8, 1         ; / 2
    imul r8, INTSIZ
    sub r9, r8        ; r9 -= (len/2) * sizeof(int)
    add r8, INTSIZ    ; r8 = &l->items[(len/2) * sizeof(int)]
    add r8, rdi

list_reverse_loop:
    mov r10, [r8]
    xchg [r9], r10
    mov [r8], r10

    sub r8, INTSIZ
    add r9, INTSIZ

    cmp r8, rdi
    jne list_reverse_loop

    mov rax, r9
    sub rax, rdi

    pop r10
    pop r9
    pop r8
    ret


;; struct list *list_copy(struct list *l)
list_copy:
    push rdi
    push rsi
    push rcx

    ;; struct list *new = malloc((l->len + 1)*8)
    mov rsi, rdi
    mov rdi, [rsi]
    inc rdi
    imul rdi, INTSIZ
    push rsi
    call malloc
    pop rsi

    test rax, rax
    jz list_copy_ret

    ;; memcpy(new, l, l->len+1)
    mov rcx, [rsi]
    inc rcx
    mov rdi, rax
    rep movsq

list_copy_ret:
    pop rcx
    pop rsi
    pop rdi
    ret


;; IMPLEMENT MULTIPLE ARGS LATER
;; struct list *list_iterate(struct list *l, void func())
list_iterate:
    push rbp
    mov rbp, rsp
    sub rsp, 0x8

    push rdi
    push rsi
    push r8

    call list
    test rax, rax
    jz list_iterate_ret
    mov [rbp], rax

    mov r8, rdi
    mov rcx, [rdi]

list_iterate_loop:
    add r8, 8
    mov rdi, [r8]
    call rsi

    mov rdi, [rbp]
    mov rsi, rax
    call list_append
    test rax, rax
    jz list_iterate_ret
    mov [rbp], rax

list_iterate_call:
    loop list_iterate_loop
    mov rax, [rbp]

list_iterate_ret:
    pop r8
    pop rsi
    pop rdi

    leave
    ret


; int list_count(struct list *l, int value)
list_count:
    push rdi
    push rcx
    push r8

    xor r8, r8
    mov rcx, [rdi]

list_count_loop:
    cmp [rdi + rcx*8], rsi
    jne list_count_not_matching
    inc r8

list_count_not_matching:
    loop list_count_loop

    mov rax, r8

    pop r8
    pop rcx
    pop rdi
    ret


;; struct list *list_join(struct list *l1, struct list *l2)
list_join:
    push rdi
    push rsi
    push rcx
    push r8
    push r9

    mov r8, rsi
    mov rcx, [rsi]

list_join_loop:
    add r8, 8
    mov rsi, [r8]

    push rcx
    call list_append
    pop rcx

    test rax, rax
    jz list_join_ret
    mov rdi, rax

    loop list_join_loop
    mov rax, rdi

list_join_ret:
    pop r9
    pop r8
    pop rcx
    pop rsi
    pop rdi
    ret


;; int mod_by_idx(struct list *l, int idx, int val)
mod_by_idx:
    xor rax, rax
    cmp rsi, [rdi]
    jg mod_by_idx_fail

    mov [rdi + 1 + rsi*INTSIZ], rdx
    jmp mod_by_idx_ret                      ; ret 0

mod_by_idx_fail:
    dec rax                                 ; ret -1

mod_by_idx_ret:
    ret


;; int get_by_idx(struct list *l, int idx)
get_by_idx:
    mov rax, [rdi + 1 + rsi*INTSIZ]
    ret

;; void list_print(struct list *l)
list_print:
    push rdi
    push rsi
    push rdx
    push rcx
    push r8
    push r9

    mov rsi, rdi
    mov rcx, [rsi]

    push rbp
    mov rbp, rsp
    mov r8, rcx
    inc r8
    imul r8, 20
    sub rsp, 0x8*2
    sub rsp, r8

    ;; r8 = buf[l->len]
    mov r8, rsp

    mov byte [r8], '['
    inc r8

list_print_loop:
    add rsi, 8
    mov rdi, [rsi]
    call itoa
    test rax, rax
    jz list_print_ret

    test rcx, rcx
    jz list_print_loop_end

list_print_copy:
    mov dl, [rax]
    test dl, dl
    jz list_print_copy_end
    mov [r8], dl
    inc r8
    inc rax
    jmp list_print_copy

list_print_copy_end:
    dec rcx
    test rcx, rcx
    jz list_print_loop_end
    mov byte [r8], ','
    inc r8
    mov byte [r8], ' '
    inc r8
    jmp list_print_loop

list_print_loop_end:
    mov byte [r8], ']'
    inc r8
    mov byte [r8], 0xa
    inc r8

    xor rdi, rdi
    xor rax, rax
    inc di
    mov rsi, rsp
    mov rdx, r8
    sub rdx, rsi
    mov al, SYS_WRITE
    syscall
    jmp list_print_ret

list_print_ret:
    leave

    pop r9
    pop r8
    pop rcx
    pop rdx
    pop rsi
    pop rdi
    ret


;; void list_send_fd_L(struct list *l, int fd_L)
list_send_fd_L:
    push rdi
    push rsi
    push rdx
    push rcx
    push r8
    push r9
    push r11

    mov r11, rsi
    mov rsi, rdi
    mov rcx, [rsi]

    push rbp
    mov rbp, rsp
    mov r8, rcx
    inc r8
    imul r8, 20
    sub rsp, 0x8*2
    sub rsp, r8

    ;; r8 = buf[l->len]
    mov r8, rsp

list_send_fd_L_loop:
    add rsi, 8
    mov rdi, [rsi]
    call itoa
    test rax, rax
    jz list_send_fd_L_ret

    test rcx, rcx
    jz list_send_fd_L_loop_end

    ; copy itoa result to stack
list_send_fd_L_copy:
    mov dl, [rax]
    test dl, dl
    jz list_send_fd_L_copy_end
    mov [r8], dl
    inc r8
    inc rax
    jmp list_send_fd_L_copy

list_send_fd_L_copy_end:
    dec rcx
    test rcx, rcx
    jz list_send_fd_L_loop_end
    mov byte [r8], ' '
    inc r8
    jmp list_send_fd_L_loop

list_send_fd_L_loop_end:
    mov byte [r8], 0xa
    inc r8

    mov rdi, r11
    xor rax, rax
    mov rsi, rsp
    mov rdx, r8
    sub rdx, rsi
    mov al, SYS_WRITE
    syscall
    jmp list_send_fd_L_ret

list_send_fd_L_ret:
    leave

    pop r11
    pop r9
    pop r8
    pop rcx
    pop rdx
    pop rsi
    pop rdi
    ret


;; void list_send_fd_R(struct list *l, int fd_R)
list_send_fd_R:
    push rdi
    push rsi
    push rdx
    push rcx
    push r8
    push r9
    push r11

    mov r11, rsi
    mov rsi, rdi
    mov rcx, [rsi]

    push rbp
    mov rbp, rsp
    mov r8, rcx
    inc r8
    imul r8, 20
    sub rsp, 0x8*2
    sub rsp, r8

    ;; r8 = buf[l->len]
    mov r8, rsp

    ;; rsi = &l->items[l->len]
    mov r9, rcx
    inc r9
    imul r9, 8
    add rsi, r9

list_send_fd_R_loop:
    sub rsi, 8
    mov rdi, [rsi]
    call itoa
    test rax, rax
    jz list_send_fd_R_ret

    test rcx, rcx
    jz list_send_fd_R_loop_end

    ; copy itoa result to stack
list_send_fd_R_copy:
    mov dl, [rax]
    test dl, dl
    jz list_send_fd_R_copy_end
    mov [r8], dl
    inc r8
    inc rax
    jmp list_send_fd_R_copy

list_send_fd_R_copy_end:
    dec rcx
    test rcx, rcx
    jz list_send_fd_R_loop_end
    mov byte [r8], ' '
    inc r8
    jmp list_send_fd_R_loop

list_send_fd_R_loop_end:
    mov byte [r8], 0xa
    inc r8

    mov rdi, r11
    xor rax, rax
    mov rsi, rsp
    mov rdx, r8
    sub rdx, rsi
    mov al, SYS_WRITE
    syscall
    jmp list_send_fd_R_ret

list_send_fd_R_ret:
    leave

    pop r11
    pop r9
    pop r8
    pop rcx
    pop rdx
    pop rsi
    pop rdi
    ret


;; void free_list(struct list *l)
free_list:
    call free
    ret
