package compiler

import (
	"lisp/lisp/go/cli"
	"lisp/lisp/go/types"
)

func Compile(source string) {
	tokens := lexer(source)
	cli.Debug("tokens", tokens)
}

func lexer(source string) types.Tokens {
	source += "\n"
	current := 0
	var tokens types.Tokens
	for current < len([]rune(source)) {
		char := string([]rune(source)[current])
		cli.Debug("token", char)
		if char == "(" {
			tokens.Append("paren", "(")
			goto next
		}
		if char == ")" {
			tokens.Append("paren", ")")
			goto next
		}
		if isNumber(char) {
			value := ""
			for isNumber(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}
			tokens.Append("number", value)
			goto next
		}
		if isAlpha(char) {
			value := ""
			for isAlpha(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}
			tokens.Append("name", value)
			goto next
		}
		break
	next:
		current++
		continue
	}
	return tokens
}

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isAlpha(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}

func ast()        {}
func generation() {}
