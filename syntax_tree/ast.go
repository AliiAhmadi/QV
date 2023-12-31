package ast

type Node interface {
	TokenLiteral() string
}

type Command interface {
	Node
	commandNode()
}

type Query struct {
	Command *Command
}

func (query *Query) TokenLiteral() string {
	if query.Command != nil {
		return query.Command.TokenLiteral()
	} else {
		return ""
	}
}
