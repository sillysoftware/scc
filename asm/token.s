section .data
  newline db 0xA
; -------------------------
; "EXIT" STATEMENT
; -------------------------
exit:
  mov rax, 60
  mov rdi, r8
  syscall
; -------------------------
; "PRINT" STATEMENT
; -------------------------
print:
  mov rax, 1
  mov rdi, 1
  mov rsi, r8
  mov rdx, 8
  syscall
; -------------------------
; "PRINTLN" STATEMENT
; -------------------------
println:
  mov rax, 1
  mov rdi, 1
  mov rsi, r8
  mov rdx, 8
  syscall
  mov rax, 1
  mov rdi, 1
  mov rsi, 0xA
  mov rdx, 1
  syscall
