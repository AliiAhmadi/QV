package parser

import (
	"QV/lexer"
	ast "QV/syntax_tree"
	"QV/token"
	"testing"
)

func TestNameOfTableAndCreateTable(t *testing.T) {
	t.Parallel()

	inputs := []string{
		`
		CREATE TABLE test12 ();
		`,
		`
		CREATE TABLE hello ( cl1 INT );
		`,
		`
		CREATE TABLE _name_start_with_under_line (
			is_ok BOOLEAN,
			col2 VARCHAR(200)
		);
		`,
	}

	tests := []struct {
		expectedName string
	}{
		{"test12"},
		{"hello"},
		{"_name_start_with_under_line"},
	}

	for index, input := range inputs {
		lex := lexer.New(input)
		pars := New(lex)

		program := pars.Parse()
		if program == nil {
			t.Fatalf("inputs[%d] - parsing program failed - returned nil", index)
		}

		if program.Query.Name == nil {
			t.Fatalf("inputs[%d] - program.Query.Name is nil", index)
		}

		if !testTable(t, &program.Query, tests[index].expectedName) {
			return
		}
	}

}

func testTable(t *testing.T, query ast.Query, tableName string) bool {
	if query.TokenLiteral() != "CREATE" {
		t.Errorf("query has incorrect token literal expected=%s, got=%s", "CREATE", query.TokenLiteral())
		return false
	}

	create, ok := query.(*ast.CreateQuery)
	if !ok {
		t.Errorf("query is not a CreateQuery. got=%q", query)
		return false
	}

	if create.Name.Token.Type != token.NAME {
		t.Errorf("create name type mismatch. ecpected='%s' got='%s'", token.NAME, create.Name.Token.Type)
		return false
	}

	if create.Name.Value != tableName {
		t.Errorf("create.Name.Value not '%s'. got='%s'", tableName, create.Name.Value)
		return false
	}

	return true
}
