global _start

_start:
  call main
  call main_exit

main_exit:
  mov rax, 60
  mov rdi, 0
  syscall

main:
