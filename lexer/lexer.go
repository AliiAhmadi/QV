package lexer

import "QV/token"

// Lexer struct contain query string
// for evaluating tokens.
type Lexer struct {
	input        string
	ch           byte // current char under examination
	position     int  // current position in input
	readPosition int  // current reading position in input
}

// Create a new Lexer pointer
func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	lexer.readChar()
	return lexer
}

func (lexer *Lexer) skipWhiteSpaces() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\t' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

// Read a character from input to lexer instance.
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func newToken(tokType token.Type, ch byte) token.Token {
	return token.Token{
		Type:    tokType,
		Literal: string(ch),
	}
}

func (lexer *Lexer) readString() string {
	position := lexer.position + 1

	for {
		lexer.readChar()

		if lexer.ch == '"' || lexer.ch == 0 {
			break
		}
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readByte() string {
	position := lexer.position + 1

	lexer.readChar()
	lexer.readChar()

	return lexer.input[position : position+1]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && 'Z' >= ch) || (ch == '_')
}

func (lexer *Lexer) readName() string {
	position := lexer.position

	if isLetter(lexer.ch) {
		lexer.readChar()
	}

	for isLetter(lexer.ch) || isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.skipWhiteSpaces()

	switch lexer.ch {
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '(':
		tok = newToken(token.LPARENTHESIS, lexer.ch)
	case ')':
		tok = newToken(token.RPARENTHESIS, lexer.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lexer.readString()
	case '\'':
		tok.Type = token.BYTE
		tok.Literal = lexer.readByte()
	case 0:
		tok.Type = token.EOQ
		tok.Literal = ""
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readName()
			tok.Type = token.Lookup(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.NUMBER
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}
