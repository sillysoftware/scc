package types

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
	prog string
	operations []string
}

func (a Assembly)Reduce() {
	//reduce from [asm1, asm2, asm3] -> "asm1\nasm2\nasm3"
}
