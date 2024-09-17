// main.go defines main() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of LISP.
//
//	LISP is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package main

import (
	"lisp/lisp/cli"
	"lisp/lisp/compiler"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cli.Fatal("No files found")
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(content)
	compiler.Compile(input)
}
