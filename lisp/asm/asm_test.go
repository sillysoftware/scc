package asm_test

import (
	_ "embed"
	"testing"
)

var (
	//go:embed assembly/test.s
	testAsset string
)

func TestAssembly(t *testing.T) {
}
