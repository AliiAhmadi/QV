package ast

import "QV/token"

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Query CreateQuery
}

type CreateQuery struct {
	Token   token.Token
	Name    *Identifier
	Columns []Column
}

type Identifier struct {
	Token token.Token
	Value string
}

type Column struct {
	Token    token.Token
	Name     *Identifier
	DataType token.DataType
	MetaData token.MetaData
}

func (program *Program) TokenLiteral() string         { return program.Query.TokenLiteral() }
func (createQuery *CreateQuery) TokenLiteral() string { return createQuery.Token.Literal }
func (identifier *Identifier) TokenLiteral() string   { return identifier.Token.Literal }
func (column *Column) TokenLiteral() string           { return column.Token.Literal }
