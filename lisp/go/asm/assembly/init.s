section .data
; term .data
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
