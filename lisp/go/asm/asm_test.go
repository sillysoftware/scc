package asm_test

import (
  "testing"
  "lisp/lisp/go/asm"
)

func TestAsm(t *testing.T) {
  var assm types.Assembly
	init := asm.GenStartAsm("test.lisp")
	write := asm.GenWriteAsm(1, "mayo", 4)
	assm.Asm = append(assm.Asm, init)
	assm.Asm = append(assm.Asm, write)
	cli.Debug("asm", assm.Reduce())
}
