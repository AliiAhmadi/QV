package ast

import "QV/token"

type Node interface {
	TokenLiteral() string
}

type Command interface {
	Node
	commandNode()
}

type Query struct {
	Command Command
}

func (query *Query) TokenLiteral() string {
	if query.Command != nil {
		return query.Command.TokenLiteral()
	} else {
		return ""
	}
}

type Column struct {
	Token token.Token // NAME
	Name  *Identifier
	Type  token.Type
}

type CreateQuery struct {
	Token   token.Token // CREATE
	Name    *Identifier
	Value   Command
	Columns []Column
}

func (createQuery *CreateQuery) commandNode()         {}
func (createQuery *CreateQuery) TokenLiteral() string { return createQuery.Token.Literal }

type Identifier struct {
	Token token.Token // NAME
	Value string
}

func (identifier *Identifier) commandNode()         {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }
