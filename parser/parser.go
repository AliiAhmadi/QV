package parser

import (
	"QV/lexer"
	ast "QV/syntax_tree"
	"QV/token"
	"fmt"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:    lex,
		errors: []string{},
	}

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lex.NextToken()
}

func (parser *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Query = ast.CreateQuery{}

	for parser.currentToken.Type != token.EOQ {
		q := parser.parseQuery()

		val, ok := q.(*ast.CreateQuery)
		if !ok {
			return nil
		}

		if q != nil {
			program.Query = *val
		}

		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseQuery() ast.Query {
	switch parser.currentToken.Type {
	case token.CREATE:
		q := parser.parseCreateQuery()
		if q == nil {
			return nil
		} else {
			return q
		}
	default:
		return nil
	}
}

func (parser *Parser) parseCreateQuery() *ast.CreateQuery {
	createQuery := &ast.CreateQuery{
		Token: parser.currentToken,
	}

	if !parser.expectPeek(token.TABLE) {
		return nil
	}

	if !parser.expectPeek(token.NAME) {
		return nil
	}

	createQuery.Name = &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}

	if !parser.expectPeek(token.LPARENTHESIS) {
		return nil
	}

	for !parser.expectPeek(token.RPARENTHESIS) {
		column, err := parser.parseColumn()
		if err != nil {

		}
	}

	// skip to end of query for now
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return createQuery
}

func (parser *Parser) parseColumn() (*ast.Column, error) {

}

func (parser *Parser) curTokenIs(tok token.Type) bool {
	return parser.currentToken.Type == tok
}

func (parser *Parser) peekTokenIs(tok token.Type) bool {
	return parser.peekToken.Type == tok
}

func (parser *Parser) expectPeek(tok token.Type) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(tok)
		return false
	}
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(tok token.Type) {
	message := fmt.Sprintf("expected next token to be '%s', got '%s'", tok, parser.peekToken.Type)
	parser.errors = append(parser.errors, message)
}
