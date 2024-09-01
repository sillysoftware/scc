package asm

import _ "embed"

var (
	//go:embed assembly/exit0.s
	Exit0 string
	//go:embed assembly/init.s
	Init string
)
