// main.go defines main() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of scc.
//
//	scc is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package main

import (
	"os"
	"scc/scc/cli"
	"scc/scc/compiler"
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
