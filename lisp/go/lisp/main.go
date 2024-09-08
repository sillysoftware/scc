package main

import (
	"fmt"
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
	var asmOut types.Assembly
	asmOut.Init()
	asmOut.AppendData(asm.GenWordData("lisp"))
	asmOut.Append(asm.GenWrite(1, "lisp", 0))
	asmOut.Reduce()
	fmt.Print(asmOut.Operations)
}
