package main

import (
	"lisp/lisp/go/cli"
	"lisp/lisp/go/asm"
	"lisp/lisp/go/compiler"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// exits
		cli.Fatal("No files found")
	}
	if !strings.Contains(os.Args[1], ".lisp") || !strings.Contains(os.Args[1], ".cl") {
		cli.Fatal("Incorrect file extention")
	}
	cli.Debug("asm", asm.Exit)
	cli.Debug("asm", asm.GenAsmExit(0))
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(content)
	compiler.Compile(input)
}
