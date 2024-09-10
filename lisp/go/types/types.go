/* types.go defines Token, Tokens, .Append(), Node, Visitor, Assembly, Init(), Reduce(), Append(), AppendData() for Golang

	This file is a part of LISP.

*/
package types

import "lisp/lisp/go/asm"

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
