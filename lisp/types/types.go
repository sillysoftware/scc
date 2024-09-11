// types.go defines Token, Tokens, .Append(), Node, Visitor, Assembly, Init(), Reduce(), Append(), AppendData() for Golang
//
//	Copyright (C) 2024-2024 vx-clutch
//
//	This file is part of LISP.
//
//	LISP is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package types

import "lisp/lisp/asm"

type Token struct {
	Kind  string
	Value string
}

type Tokens []Token

func (t *Tokens) Append(kind string, value string) {
	*t = append(*t, Token{Kind: kind, Value: value})
}

type Node struct {
	Kind       string
	Value      string
	Name       string
	Callee     *Node
	Expression *Node
	Body       []Node
	Params     []Node
	Arguments  *[]Node
	Context    *[]Node
}

type Visitor map[string]func(n *Node, p Node)

type Assembly struct {
	Operations string
}

func (a *Assembly) Init() {
	a.Append(asm.Init)
}

func (a *Assembly) Reduce() {
	a.Operations = "section .data\n" + a.Operations
	a.Operations += "  ret"
}

func (a *Assembly) Append(data string) {
	a.Operations += data
}

func (a *Assembly) AppendData(data string) {
	a.Operations = data + a.Operations
}
