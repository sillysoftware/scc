package asm

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed assembly/init.s
	Init string
	//go:embed assembly/exit.s
	exit string
)

func arg(n int) string {
	return fmt.Sprintf("${%d}", n)
}

func GenAsmExit(status int) string {
	out := strings.ReplaceAll(exit, arg(0), fmt.Sprintf("%d", status))
	return out
}
