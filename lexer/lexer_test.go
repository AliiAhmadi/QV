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

func TestFullCreateTableQuery(t *testing.T) {
	t.Parallel()
	query := `
	CREATE TABLE customers (
		customerID INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		address TEXT,
		email TEXT
	  );
	`

	toks := []struct {
		name            string
		expectedType    token.Type
		expectedLiteral string
	}{
		{"create keyword", token.CREATE, "CREATE"},
		{"table keyword", token.TABLE, "TABLE"},
		{"table name", token.NAME, "customers"},
		{"left parenthesis", token.LPARENTHESIS, "("},
		{"column name", token.NAME, "customerID"},
		{"integer keyword", token.INTEGER, "INTEGER"},
		{"primary keyword", token.PRIMARY, "PRIMARY"},
		{"key keyword", token.KEY, "KEY"},
		{"comma", token.COMMA, ","},
		{"column name", token.NAME, "name"},
		{"TEXT keyword", token.TEXT, "TEXT"},
		{"NOT keyword", token.NOT, "NOT"},
		{"NULL keyword", token.NULL, "NULL"},
		{"comma", token.COMMA, ","},
		{"column name", token.NAME, "address"},
		{"TEXT keyword", token.TEXT, "TEXT"},
		{"comma", token.COMMA, ","},
		{"column name", token.NAME, "email"},
		{"TEXT keyword", token.TEXT, "TEXT"},
		{"right parenthesis", token.RPARENTHESIS, ")"},
		{"semicolon", token.SEMICOLON, ";"},
	}

	lex := New(query)

	for i, tok := range toks {
		tk := lex.NextToken()

		if tk.Literal != tok.expectedLiteral {
			t.Fatalf("tests[%d] incorrect literal. expected=%s, got=%s", i, tok.expectedLiteral, tk.Literal)
		}

		if tk.Type != tok.expectedType {
			t.Fatalf("tests[%d] incorrect type. expected=%s, got=%s", i, tok.expectedType, tk.Type)
		}
	}
}
