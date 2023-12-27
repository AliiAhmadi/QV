package lexer

import (
	"QV/token"
	"testing"
)

func TestTokens(t *testing.T) {
	t.Parallel()
	input := `CREATE TABLE hello();`

	tests := []struct {
		name            string
		expectedType    token.Type
		expectedLiteral string
	}{
		{"CREATE", token.CREATE, "CREATE"},
		{"TABLE", token.TABLE, "TABLE"},
		{"table name", token.NAME, "hello"},
		{"left parenthesis", token.LPARENTHESIS, "("},
		{"right parenthesis", token.RPARENTHESIS, ")"},
		{"semicolon", token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] incorrect type. expected=%s, got=%s", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] incorrect literal. expected=%s, got=%s", i, test.expectedLiteral, tok.Literal)
		}
	}
}
