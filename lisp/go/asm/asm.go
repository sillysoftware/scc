package asm

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed assembly/init.s
	Init string
)

func GenAsmExit(status int) string {
	return fmt.Sprintf("\tmov rax, 60\n\tmov rbi, %d\n\tsyscall\n", status)
}
