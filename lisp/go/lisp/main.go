package main

import (
	"lisp/lisp/go/cli"
	"lisp/lisp/go/compiler"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// exits
		cli.Fatal("No files found")
	}
	cli.Warn("This is an unsafe cli. It will just attempt to compile what ever argument is passed to this.")
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(content)
	compiler.Compile(input)
}
