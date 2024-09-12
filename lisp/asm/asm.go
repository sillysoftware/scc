// asm.go defines Init, GenWordData(), GenExit(), GenWrite() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of LISP.
//
//	LISP is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package asm

import (
	_ "embed"
	"fmt"
	. "strconv"
	. "strings"
)

var (
	Assembly string
	Data     string
	//go:embed assembly/data.s
	dataSection string
	//go:embed assembly/init.s
	Init string
	//go:embed assembly/exit.s
	exit string
	//go:embed assembly/print.s
	write string
)

func appendAsm(data string) {
	Assembly += data
}

func Reduce() string {
	buf := Init + Assembly + "  ret"
	return buf
}

func arg(n int) string {
	return fmt.Sprintf("${%d}", n)
}

func GenExit(status int) {
	buf := exit
	buf = ReplaceAll(buf, arg(0), Itoa(status))
	appendAsm(buf)
}

func GenWrite(desc int, word string) {
	buf := write
	buf = ReplaceAll(buf, arg(0), Itoa(desc))
}
