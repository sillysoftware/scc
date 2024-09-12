package asm_test

import (
	_ "embed"
	"lisp/lisp/asm"
	"testing"
)

var (
	//go:embed assembly/test.s
	testAsset string
)

func TestAssembly(t *testing.T) {
	asm.GenWrite(1, "lisp")
	buf := asm.Reduce()
	if buf != testAsset {
		t.Fatal("Incorrect Asm output")
	}
}
