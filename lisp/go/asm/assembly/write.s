  mov rax, 1
  mov rdi, ${0} ; File descriptor
  mov rsi, ${1} ; .data word
  mov rdx, ${2} ; .data word size
  syscall
