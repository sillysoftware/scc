package asm_test

import (
	_ "embed"
	"lisp/lisp/go/asm"
	"lisp/lisp/go/types"
	"os"
	"testing"
)

var (
	//go:embed assembly/test.s
	testAsset string
)

func TestAssembly(t *testing.T) {
	var asmOut types.Assembly
	asmOut.Init()
	asmOut.AppendData(asm.GenWordData("lisp"))
	asmOut.Append(asm.GenWrite(1, "lisp", 0))
	asmOut.Reduce()
	if asmOut.Operations != testAsset {
		os.WriteFile("asmlog", []byte(asmOut.Operations), 0755)
		t.Fatal("Incorrect assembly generated logs outputed to asmlog")
	}
}
