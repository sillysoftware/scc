package main

import (
	"lisp/lisp/go/asm"
	"lisp/lisp/go/cli"
	"lisp/lisp/go/compiler"
	"os"
)

func main() {
	cli.Warn("This is an unsafe cli. It will just attempt to compile what ever argument is passed to this.")
	cli.Debug("asm exit0 output", asm.Exit_0)
	compiler.Compile(os.Args[1])
}
