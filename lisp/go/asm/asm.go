package asm

import (
	_ "embed"
	"fmt"
	. "strconv"
	. "strings"
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

func GenWrite(desc int, word string) (string, string) {
	buf := ReplaceAll(write, arg(0), Itoa(desc))
	buf = ReplaceAll(buf, arg(1), word)
	buf = ReplaceAll(buf, arg(2), Itoa(len(word)))
	return buf, AddData(word)
}

func GenExit(status int) string {
	buf := ReplaceAll(exit, arg(0), Itoa(status))
	return buf
}

func AddData(word string) string {
	return ""
}
