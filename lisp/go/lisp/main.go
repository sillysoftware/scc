package main

import (
	"lisp/lisp/go/asm"
	"lisp/lisp/go/cli"
	"lisp/lisp/go/compiler"
	"lisp/lisp/go/types"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// exits
		cli.Fatal("No files found")
	}
	if strings.Contains(os.Args[1], ".lisp") == false || !strings.Contains(os.Args[1], ".cl") == false {
		cli.Fatal("Incorrect file extention")
	}
	test()
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(content)
	compiler.Compile(input)
}

func test() {
	var assm types.Assembly
	write := asm.GenWriteAsm(1, "mayo")
	assm.Asm = append(assm.Asm, asm.Init)
	assm.Asm = append(assm.Asm, write)
	cli.Debug("asm", assm.Reduce())
}
