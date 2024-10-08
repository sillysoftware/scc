; -------------------------
; If you want to be a cool little hacker boy, you can add new statments by adding a new macro here and in scc/tokens.def.
section .data
  newline db 0xA
; -------------------------
; "EXIT" STATEMENT
; -------------------------
%macro exit 1
	mov rax, 60
 	mov rdi, %1
  	syscall
%endmacro
; -------------------------
; "PRINT" STATEMENT
; -------------------------
%macro print 2
  mov rax, 1
  mov rdi, 1
  mov rsi, %1
  mov rdx, %2
  syscall
%endmacro
; -------------------------
; "PRINTLN" STATEMENT
; -------------------------
%macro print 2
  mov rax, 1
  mov rdi, 1
  mov rsi, %1
  mov rdx, %2
  syscall
  print newline, 1
%endmacro
