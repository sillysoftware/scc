  mov rax, 1
  mov rdi, ${0} ; File descriptor
  mov rsi, ${1} ; .bss word
  mov rdx, ${2} ; .bss word size
  syscall
