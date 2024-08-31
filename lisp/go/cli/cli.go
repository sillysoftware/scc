package cli

import (
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	errC   = "\033[091m" // HI-Red
	warnC  = "\033[095m" // Hi-Purple
	debugC = "\033[097m" // Hi-White
)

func Fatal(m string) {
	fmt.Printf("LITX: %sfatal error%s: %s", errC, reset, m)
	fmt.Println("\ncompilation terminated.")
	os.Exit(0)
}

func Error(m string) {
	fmt.Printf("LITX: %serror%s: %s\n", errC, reset, m)
}

func Warn(warning string) {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, warning)
}

func Debug(tag string, dbm any) {
	fmt.Printf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
}
