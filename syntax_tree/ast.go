package ast

import (
	"QV/token"
	"bytes"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Query interface {
	Node
	QueryNode()
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
	MetaData []token.MetaData
}

func (program *Program) TokenLiteral() string { return program.Query.TokenLiteral() }

func (program *Program) QueryNode() {}

func (program *Program) String() string { return program.Query.String() }

func (createQuery *CreateQuery) TokenLiteral() string { return createQuery.Token.Literal }

func (createQuery *CreateQuery) QueryNode() {}

func (createQuery *CreateQuery) String() string {
	var output bytes.Buffer

	output.WriteString(createQuery.Token.Literal + " " + token.TABLE + " ")
	output.WriteString(createQuery.Name.String())
	output.WriteString("( ")

	columns := []string{}

	for _, col := range createQuery.Columns {
		columns = append(columns, col.String())
	}

	output.WriteString(strings.Join(columns, ", "))
	output.WriteString(" )")

	return output.String()
}

func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }

func (identifier *Identifier) QueryNode() {}

func (identifier *Identifier) String() string { return identifier.Value }

func (column *Column) TokenLiteral() string { return column.Token.Literal }

func (column *Column) QueryNode() {}

func (column *Column) String() string {
	var output bytes.Buffer

	values := []string{}

	values = append(values, column.Name.String())
	values = append(values, column.DataType.String())

	for _, meta := range column.MetaData {
		values = append(values, meta.String())
	}

	output.WriteString(strings.Join(values, " "))

	return output.String()
}
