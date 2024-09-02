package compiler

import (
	"fmt"
	"lisp/lisp/go/cli"
	"lisp/lisp/go/types"
)

func Compile(source string) {
	tokens := lexer(source)
	ast := parser(tokens)
}

func lexer(source string) types.Tokens {
	current := 0
	source += "\n"
	var tokens types.Tokens
	for current < len([]rune(source)) {
		char := string([]rune(source)[current])
		if char == "(" {
			tokens.Append("paren", "(")
			goto next
		}
		if char == ")" {
			tokens.Append("paren", ")")
			goto next
		}
		if char == " " {
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
			continue
		}
		if isAlpha(char) {
			value := ""
			for isAlpha(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}
			tokens.Append("ident", value)
			continue
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

var pc int
var pt types.Tokens

type ast types.Node

func parser(tokens types.Tokens) ast {
	pc = 0
	pt = tokens
	ast := ast{
		Kind: "prog",
		Body: []types.Node{},
	}
	for pc < len(pt) {
		ast.Body = append(ast.Body, walk())
	}
	return ast
}

func walk() types.Node {
	token := pt[pc]
	if token.Kind == "number" {
		pc++
		return types.Node{Kind: "NumberLiteral", Value: token.Value}
	}
	if token.Kind == "paren" && token.Value == "(" {
		pc++
		token = pt[pc]
		n := types.Node{
			Kind:   "CallExpression",
			Name:   token.Value,
			Params: []types.Node{},
		}
		pc++
		token = pt[pc]
		for token.Kind != "paren" || (token.Kind == "paren" && token.Value != ")") {
			n.Params = append(n.Params, walk())
			token = pt[pc]
		}
		pc++
		return n
	}
	cli.Fatal(fmt.Sprintf("Unexpected token: %s", token.Value))
	return types.Node{}
}

func traverser(a ast, v types.Visitor) {
}

func traverseArray(a []types.Node, p types.Node, v types.Visitor) {
	for _, child := range a {
		traverseNode(child, p, v)
	}
}

func traverseNode(n, p types.Node, v types.Visitor) {
	for k, va := range v {
		if k == n.Kind {
			va(&n, p)
		}
	}
	switch n.Kind {
	case "prog":
		traverseArray(n.Body, n, v)
		break
	case "CallExpression":
		traverseArray(n.Params, n, v)
		break
	case "NumberLiteral":
		break
	default:
		cli.Fatal(fmt.Sprintf("Unexpected node: %s", n.Kind))
	}
}

func transformer(a ast) ast {
	nast := ast{
		Kind: "Program",
		Body: []types.Node{},
	}
	a.Context = &nast.Body
	return nast
}
