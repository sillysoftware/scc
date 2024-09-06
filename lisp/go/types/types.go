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
	Operations []string
}

func (a *Assembly) Append(data string) {
	a.Operations = append(a.Operations, data)
}

func (a Assembly) Reduce() string {
	buf := ""
	for _, opt := range a.Operations {
		buf += "\n" + opt
	}
	buf += "  ret"
	return buf
}

func (a *Assembly) Init() {
	a.Append(asm.Init)
}

func (a *Assembly) Data(data string) {
	init := a.Operations[0]
	_ = init
	// add data
}
