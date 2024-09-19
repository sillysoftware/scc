// cli.go defines Fatal(), Error(), Warn(), Debug() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of SCC.
//
//	SCC is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package cli

import (
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	errC   = "\033[1;91m" // Bold High Intensity Red
	warnC  = "\033[1;95m" // Bold High Intensity Purple
	debugC = "\033[1;97m" // Bold High Intensity White

)

func Fatal(m any) {
	fmt.Printf("SCC: %sfatal error%s: %s", errC, reset, m)
	fmt.Println("\ncompilation terminated.")
	os.Exit(0)
}

func Error(m any) string {
	fmt.Printf("SCC: %serror%s: %s\n", errC, reset, m)
	return fmt.Sprintf("SCC: %serror%s: %s\n", errC, reset, m)
}

func Warn(warning any) string {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, warning)
	return fmt.Sprintf("SCC: %serror%s: %s\n", errC, reset, warning)
}

func Debug(tag string, dbm any) string {
	fmt.Printf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
	return fmt.Sprintf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
}
