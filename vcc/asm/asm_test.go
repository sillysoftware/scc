package asm_test

import (
	_ "embed"
	"testing"
	"vcc/vcc/asm"
)

var (
	//go:embed assembly/x86-64/test.s
	testAsset string
)

func TestAssembly(t *testing.T) {
	asm.GenWrite(1, "vcc")
	buf := asm.Reduce()
	if buf != testAsset {
		t.Fatal("Incorrect Asm output")
	}
}
