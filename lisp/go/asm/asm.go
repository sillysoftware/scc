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
	Exit string
)

func arg(n int) string {
	return fmt.Sprintf("${%d}", n)
}

func GenAsmExit(stat int) string {
	status := string(stat)
	return strings.ReplaceAll(Exit, arg(0), status)
}
