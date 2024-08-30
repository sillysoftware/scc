package types

type Token struct {
	kind  string
	value string
}

type Node struct {
	kind       string
	value      string
	name       string
	callee     *Node
	expression *Node
	body       []Node
	params     []Node
	arguments  *[]Node
	context    *[]Node
}

type Ast Node

type Visitor map[string]func(n *Node, p Node)
