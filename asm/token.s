section .data
  newline db 0xA
; -------------------------
; "EXIT" STATEMENT
; -------------------------
%macro exit
	; exit
%end
; -------------------------
; "PRINT" STATEMENT
; -------------------------
print:
  mov rax, 1
  mov rdi, 1
  mov rsi, r8
  mov rdx, 8
  syscall
  ret
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
  ret
; -------------------------
; "ADD" STATEMENT
; -------------------------
add:
  xor r8, r8
  xor r9, r9
  add r8, r9
  mov rax, r8
