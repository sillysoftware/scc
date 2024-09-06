  .file "${0}"
section .text
  global _start

_start:
  call main
  jmp main_exit

main_exit:
  mov eax, 1
  xor ebx, ebx
  int 0x80

main:
