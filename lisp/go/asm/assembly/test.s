section .data
  lisp0 db 'lisp'
  lisp0l equ $-lisp0
section .text
  global _start

_start:
  call main
  jmp main_exit

main_exit:
  mov rax, 60
  mov rdi, 0
  syscall

main:
  mov rax, 1
  mov rdi, 1
  mov rsi, lisp0
  mov rdx, lisp0l
  syscall
  ret