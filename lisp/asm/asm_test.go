package asm_test

import (
	_ "embed"
	"lisp/lisp/asm"
	"lisp/lisp/cli"
	"testing"
)

var (
	//go:embed assembly/test.s
	testAsset string
)

func TestAssembly(t *testing.T) {
	asm.GenWrite(1, "lisp")
	buf := asm.Reduce()
	_ = buf
}
