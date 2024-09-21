// compiler.go defines Compile() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of scc.
//
//	scc is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package compiler

import (
	"fmt"
	"runtime"
	"strings"
	// "scc/scc/asm"
	"scc/scc/cli"
	"scc/scc/types"
)

func Compile(source string) string {
	return codeGenerator(types.Node(transformer(parser(lexer(source)))))
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
	traverseNode(types.Node(a), types.Node{}, v)
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
		Kind: "prog",
		Body: []types.Node{},
	}
	a.Context = &nast.Body
	traverser(a, map[string]func(n *types.Node, p types.Node){
		"NumberLiteral": func(n *types.Node, p types.Node) {
			*p.Context = append(*p.Context, types.Node{
				Kind:  "NumberLiteral",
				Value: n.Value,
			})
		},
		"CallExpression": func(n *types.Node, p types.Node) {
			e := types.Node{
				Kind: "CallExpression",
				Callee: &types.Node{
					Kind: "Identifier",
					Name: n.Name,
				},
				Arguments: new([]types.Node),
			}
			n.Context = e.Arguments
			if p.Kind != "CallExpression" {
				es := types.Node{
					Kind:       "ExpressionStatement",
					Expression: &e,
				}
				*p.Context = append(*p.Context, es)
			} else {
				*p.Context = append(*p.Context, e)
			}
		},
	})
	return nast
}

func codeGenerator(n types.Node) string {
	if runtime.GOOS != "linux" {
		cli.Fatal("Unsupported OS Detected: For an up-to-date support list go to the docs directory")
	}
	switch n.Kind {
	case "Program":
		var r []string
		for _, no := range n.Body {
			r = append(r, codeGenerator(no))
		}
		return strings.Join(r, "\n")
	case "ExperessionStatement":
		return codeGenerator(*n.Expression) + ";"
	case "CallExpression":
		var ra []string
		c := codeGenerator(*n.Callee)
		for _, no := range *n.Arguments {
			ra = append(ra, codeGenerator(no))
		}
		r := strings.Join(ra, ", ")
		return c + "(" + r + ")"
	case "Identifier":
		return n.Name
	case "NumberLiteral":
		return n.Value
	default:
		cli.Fatal(n)
		return ""
	}
}
