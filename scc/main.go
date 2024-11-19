package main

import (
	"fmt"
	"os"
)

type sexpr struct {
	atoms []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ftl: compile time error: not enough arguments provided.")
		os.Exit(1)
	}

	path := os.Args[1]

	s, e := os.ReadFile(path)
	if e != nil {
		fmt.Println("Error reading file: ", e)
		os.Exit(0)
	}
	expr := lexer(string(s))
	seprint(expr[0])
}

func seprint(s sexpr) {
	fmt.Print("(")
	for _, atom := range s.atoms {
		fmt.Print(" ", atom)
	}
	fmt.Println(" )")
}

func lexer(s string) []sexpr {
	sexpers := []sexpr{}
	current := 0
	s += "\n"
	var sxprs []sexpr
	for current < len([]rune(s)) {
		char := string([]rune(s)[current])
		if char == "(" {
			var list []string
			while char != ")" {
				_ = 3
			}
		}
		if char == " " {
			goto next
		}
		next:
			current++
			continue
	}
	return sexpers
}
