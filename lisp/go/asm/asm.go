package asm

import (
	_ "embed"
	"fmt"
	. "strings"
	. "strconv"
)

var (
	//go:embed assembly/init.s
	Init string
	//go:embed assembly/exit.s
	exit string
	//go:embed assembly/write.s
	write string
)

func arg(n int) string {
	return fmt.Sprintf("${%d}", n)
}

func GenWriteAsm(desc int, word string, size int) string {
	buf := ReplaceAll(write, arg(0), Itoa(desc))
	buf = ReplaceAll(buf, arg(1), word)
	buf = ReplaceAll(buf, arg(2), Itoa(size))
	return buf
}

func GenAsmExit(status int) string {
	buf := ReplaceAll(exit, arg(0), Itoa(status))
	return buf
}
