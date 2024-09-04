package asm

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed assembly/init.s
	Init string
	//go:embed assembly/exit.s
	exit string
)

func GenAsmExit(status int) string {
	return strings.Replace(exit, "$", string(status))
}
