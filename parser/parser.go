package parser

import (
	"QV/lexer"
	ast "QV/syntax_tree"
	"QV/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex: lex,
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
	return nil
}
