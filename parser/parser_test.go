package parser

import (
	"QV/lexer"
	ast "QV/syntax_tree"
	"QV/token"
	"testing"
)

func checkParseErrors(t *testing.T, parser *Parser) {
	errs := parser.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("parse error exist, count: %d", len(errs))
	for _, message := range errs {
		t.Errorf("parse error: %q", message)
	}
	t.FailNow()
}

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
		`
		CREATE TABLE usernames (
			username TEXT PRIMARY KEY,
			password TEXT NOT NULL
		);
		`,
		`
		CREATE TABLE products (
			productID INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			description TEXT,
			category TEXT
		);
		`,
		`
		CREATE TABLE orders_ (
			orderID INTEGER PRIMARY KEY,
			customerID INTEGER,
			productID INTEGER,
			quantity INTEGER NOT NULL,
			price REAL NOT NULL
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS test_test_test ();
		`,
		`
		CREATE TABLE Customers (
			First_Name varchar(50) NOT NULL,
			Last_Name varchar(50) NOT NULL,
			City varchar(50) NOT NULL,
			Email varchar(100) NOT NULL,
			Phone_Number varchar(20) NOT NULL,
			Registration_Date date NOT NULL
			);
		`,

		`
		CREATE TABLE Customers (
			First_Name varchar(50) NOT NULL,
			Last_Name varchar(50) NOT NULL,
			Email varchar(100) NOT NULL,
			Phone_Number varchar(20) NOT NULL,
			CONSTRAINT PK_Customer PRIMARY KEY (Last_Name, Email, Phone_Number)
			);
		`,
	}

	tests := []struct {
		expectedName string
	}{
		{"test12"},
		{"hello"},
		{"_name_start_with_under_line"},
		{"usernames"},
		{"products"},
		{"orders_"},
		{"test_test_test"},
		{"Customers"},
		{"Customers"},
	}

	for index, input := range inputs {
		lex := lexer.New(input)
		pars := New(lex)
		program := pars.Parse()

		checkParseErrors(t, pars)

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
