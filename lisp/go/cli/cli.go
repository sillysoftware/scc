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
	fmt.Printf("LISP: %sfatal error%s: %s", errC, reset, m)
	fmt.Println("\ncompilation terminated.")
	os.Exit(0)
}

func Error(m any) string {
	fmt.Printf("LISP: %serror%s: %s\n", errC, reset, m)
	return fmt.Sprintf("LISP: %serror%s: %s\n", errC, reset, m)
}

func Warn(warning any) string {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, warning)
	return fmt.Sprintf("LISP: %serror%s: %s\n", errC, reset, warning)
}

func Debug(tag string, dbm any) string {
	fmt.Printf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
	return fmt.Sprintf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
}
