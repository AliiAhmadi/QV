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

func TestCreatingTableQueries(t *testing.T) {
	t.Parallel()

	inputs := []string{
		`CREATE TABLE usernames (
			username TEXT PRIMARY KEY,
			password TEXT NOT NULL
		  );
		`,
		`CREATE TABLE products (
			productID INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			description TEXT,
			category TEXT
		  );	  
		`,
		`CREATE TABLE orders (
			orderID INTEGER PRIMARY KEY,
			customerID INTEGER,
			productID INTEGER,
			quantity INTEGER NOT NULL,
			price REAL NOT NULL
		  );
		`,
		`
		CREATE hello();
		`,
		`SELECT select_list FROM table ORDER BY column_1 ASC, column_2 DESC;
		`,
		`
		SELECT
			name,
			milliseconds, 
        	albumid
		FROM
			tracks;
		`,
		`
		SELECT
		name,
		milliseconds, 
		albumid
		FROM
		tracks
		ORDER BY
		albumid ASC;
		`,
		`
		SELECT * FROM locations;
		`,
		`
		SELECT region_name AS "Name of the Region" FROM regions;
		`,
		`
		SELECT * FROM locations LIMIT 10;
		`,
		`
		SELECT * FROM locations LIMIT 4,5;
		`,
		`
		SELECT * FROM locations LIMIT 5 OFFSET 4;
		`,
		`
		SELECT * FROM locations WHERE country_id='CA';
		`,
		`
		SELECT * FROM agents WHERE commission<15;
		`,
		`
		SELECT agent_code, agent_name, commission
		FROM agents 
		WHERE commission>12000;
		`,
		`
		SELECT employee_id, first_name, last_name,job_id, manager_id,department_id 
		FROM employees
		WHERE department_id=80;
		`,
		`
		SELECT employee_id, first_name, last_name,job_id, manager_id,department_id 
		FROM employees
		WHERE department_id=80 
		AND manager_id=100;
		`,
		`
		SELECT employee_id, first_name, last_name,job_id, manager_id,department_id 
		FROM employees
		WHERE job_id='SH_CLERK' 
		OR department_id=80 
		AND manager_id=147;
		`,
		`
		SELECT employee_id, first_name, last_name,job_id, manager_id,department_id 
		FROM employees
		WHERE job_id='SH_CLERK' 
		OR (department_id=80 AND manager_id=147);
		`,
		`
		SELECT DISTINCT department_id, manager_id
		FROM employees;
		`,
		`
		SELECT employee_id, first_name, last_name,job_id, salary,department_id 
		FROM employees
		WHERE department_id<50 
		ORDER BY salary;
		`,

		`
		SELECT employee_id, first_name, last_name,job_id, salary,department_id 
		FROM employees
		WHERE department_id<50 
		ORDER BY salary DESC;
		`,

		`
		SELECT employee_id, first_name, last_name,job_id, salary,department_id 
		FROM employees
		WHERE department_id<50 
		ORDER BY department_id,salary DESC;
		`,

		`
		SELECT DISTINCT department_id, manager_id,job_id
		FROM employees ORDER BY department_id;
		`,

		`
		SELECT job_id AS Designation ,"Total Salary" 
		FROM employees
		WHERE department_id<80
		GROUP BY job_id
		HAVING salery>10000
		ORDER BY salery;
		`,
	}

	tests := [][]struct {
		name            string
		expectedType    token.Type
		expectedLiteral string
	}{
		// CREATE TABLE usernames (
		// 	username TEXT PRIMARY KEY,
		// 	password TEXT NOT NULL
		//   );
		{
			{"create keyword", token.CREATE, "CREATE"},
			{"table keyword", token.TABLE, "TABLE"},
			{"table name", token.NAME, "usernames"},
			{"left pranthesis", token.LPARENTHESIS, "("},
			{"column name", token.NAME, "username"},
			{"TEXT type", token.TEXT, "TEXT"},
			{"PRIMARY", token.PRIMARY, "PRIMARY"},
			{"KEY", token.KEY, "KEY"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "password"},
			{"TEXT type", token.TEXT, "TEXT"},
			{"NOT", token.NOT, "NOT"},
			{"NULL", token.NULL, "NULL"},
			{"right pranthesis", token.RPARENTHESIS, ")"},
			{"semicolon", token.SEMICOLON, ";"},
		},
		// CREATE TABLE products (
		// 	productID INTEGER PRIMARY KEY,
		// 	name TEXT NOT NULL UNIQUE,
		// 	description TEXT,
		// 	category TEXT
		//   );
		{
			{"CREATE keyword", token.CREATE, "CREATE"},
			{"TABLE keyword", token.TABLE, "TABLE"},
			{"table name", token.NAME, "products"},
			{"left pranthesis", token.LPARENTHESIS, "("},
			{"column name", token.NAME, "productID"},
			{"INTEGER", token.INTEGER, "INTEGER"},
			{"PRIMARY", token.PRIMARY, "PRIMARY"},
			{"KEY", token.KEY, "KEY"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "name"},
			{"TEXT type", token.TEXT, "TEXT"},
			{"NOT", token.NOT, "NOT"},
			{"NULL", token.NULL, "NULL"},
			{"UNIQUE", token.UNIQUE, "UNIQUE"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "description"},
			{"TEXT type", token.TEXT, "TEXT"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "category"},
			{"TEXT type", token.TEXT, "TEXT"},
			{"right pranthesis", token.RPARENTHESIS, ")"},
			{"semicolon", token.SEMICOLON, ";"},
		},
		// CREATE TABLE orders (
		// 	orderID INTEGER PRIMARY KEY,
		// 	customerID INTEGER,
		// 	productID INTEGER,
		// 	quantity INTEGER NOT NULL,
		// 	price REAL NOT NULL
		//   );
		{
			{"CREATE keyword", token.CREATE, "CREATE"},
			{"TABLE keyword", token.TABLE, "TABLE"},
			{"table name", token.NAME, "orders"},
			{"left pranthesis", token.LPARENTHESIS, "("},
			{"column name", token.NAME, "orderID"},
			{"INTEGER", token.INTEGER, "INTEGER"},
			{"PRIMARY", token.PRIMARY, "PRIMARY"},
			{"KEY", token.KEY, "KEY"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "customerID"},
			{"INTEGER", token.INTEGER, "INTEGER"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "productID"},
			{"INTEGER", token.INTEGER, "INTEGER"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "quantity"},
			{"INTEGER", token.INTEGER, "INTEGER"},
			{"NOT", token.NOT, "NOT"},
			{"NULL", token.NULL, "NULL"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "price"},
			{"REAL", token.REAL, "REAL"},
			{"NOT", token.NOT, "NOT"},
			{"NULL", token.NULL, "NULL"},
			{"right pranthesis", token.RPARENTHESIS, ")"},
			{"semicolon", token.SEMICOLON, ";"},
		},
		{
			{"CREATE keyword", token.CREATE, "CREATE"},
			{"table name", token.NAME, "hello"},
			{"left pranthesis", token.LPARENTHESIS, "("},
			{"right pranthesis", token.RPARENTHESIS, ")"},
			{"semicolon", token.SEMICOLON, ";"},
		},
		//SELECT select_list FROM table ORDER BY column_1 ASC, column_2 DESC;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "select_list"},
			{"FROM", token.FROM, "FROM"},
			{"table", token.NAME, "table"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "column_1"},
			{"ASC", token.ASC, "ASC"},
			{"comma", token.COMMA, ","},
			{"column_2", token.NAME, "column_2"},
			{"DESC keyword", token.DESC, "DESC"},
			{"semocolon", token.SEMICOLON, ";"},
		},

		// SELECT
		// 	name,
		// 	milliseconds,
		// 	albumid
		// FROM
		// 	tracks;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "milliseconds"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "albumid"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "tracks"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT
		// name,
		// milliseconds,
		// albumid
		// FROM
		// tracks
		// ORDER BY
		// albumid ASC;

		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "milliseconds"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "albumid"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "tracks"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "albumid"},
			{"ASC", token.ASC, "ASC"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM locations;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "locations"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT region_name AS "Name of the Region" FROM regions;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "region_name"},
			{"AS", token.AS, "AS"},
			{"string", token.STRING, "Name of the Region"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "regions"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM locations LIMIT 10;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "locations"},
			{"LIMIT", token.LIMIT, "LIMIT"},
			{"number 10", token.NUMBER, "10"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM locations LIMIT 4,5;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "locations"},
			{"LIMIT", token.LIMIT, "LIMIT"},
			{"number 4", token.NUMBER, "4"},
			{"comma", token.COMMA, ","},
			{"number 5", token.NUMBER, "5"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM locations LIMIT 5 OFFSET 4;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "locations"},
			{"LIMIT", token.LIMIT, "LIMIT"},
			{"number 5", token.NUMBER, "5"},
			{"OFFSET", token.OFFSET, "OFFSET"},
			{"number 4", token.NUMBER, "4"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM locations WHERE country_id='CA';
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "locations"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "country_id"},
			{"equal sign", token.EQUALITY, "="},
			{"string", token.STRING, "CA"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT * FROM agents WHERE commission<15;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"STAR", token.STAR, "*"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "agents"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "commission"},
			{"LESSTHAN", token.LESSTHAN, "<"},
			{"NUMBER", token.NUMBER, "15"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT agent_code, agent_name, commission
		// FROM agents
		// WHERE commission>12000;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "agent_code"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "agent_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "commission"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "agents"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "commission"},
			{"GREATERTHAN", token.GREATERTHAN, ">"},
			{"NUMBER", token.NUMBER, "12000"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, manager_id,department_id
		// FROM employees
		// WHERE department_id=80;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "80"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, manager_id,department_id
		// FROM employees
		// WHERE department_id=80
		// AND manager_id=100;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "80"},
			{"AND", token.AND, "AND"},
			{"column name", token.NAME, "manager_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "100"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, manager_id,department_id
		// FROM employees
		// WHERE job_id='SH_CLERK'
		// OR department_id=80
		// AND manager_id=147;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "job_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"STR", token.STRING, "SH_CLERK"},
			{"OR", token.OR, "OR"},
			{"column name", token.NAME, "department_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "80"},
			{"AND", token.AND, "AND"},
			{"column name", token.NAME, "manager_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "147"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, manager_id,department_id
		// FROM employees
		// WHERE job_id='SH_CLERK'
		// OR (department_id=80 AND manager_id=147);
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "job_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"STR", token.STRING, "SH_CLERK"},
			{"OR", token.OR, "OR"},
			{"left parenthesis", token.LPARENTHESIS, "("},
			{"column name", token.NAME, "department_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "80"},
			{"AND", token.AND, "AND"},
			{"column name", token.NAME, "manager_id"},
			{"EQUALITY", token.EQUALITY, "="},
			{"NUMBER", token.NUMBER, "147"},
			{"right parenthesis", token.RPARENTHESIS, ")"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT DISTINCT department_id, manager_id
		// FROM employees;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"DISTINCT", token.DISTINCT, "DISTINCT"},
			{"column name", token.NAME, "department_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, salary,department_id
		// FROM employees
		// WHERE department_id<50
		// ORDER BY salary;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "salary"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"LESSTHAN", token.LESSTHAN, "<"},
			{"NUMBER", token.NUMBER, "50"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "salary"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, salary,department_id
		// FROM employees
		// WHERE department_id<50
		// ORDER BY salary DESC;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "salary"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"lessthan", token.LESSTHAN, "<"},
			{"NUMBER", token.NUMBER, "50"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "salary"},
			{"DESC", token.DESC, "DESC"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT employee_id, first_name, last_name,job_id, salary,department_id
		// FROM employees
		// WHERE department_id<50
		// ORDER BY department_id,salary DESC;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "employee_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "first_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "last_name"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "salary"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "department_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"lessthan", token.LESSTHAN, "<"},
			{"NUMBER", token.NUMBER, "50"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "department_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "salary"},
			{"DESC", token.DESC, "DESC"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT DISTINCT department_id, manager_id,job_id
		// FROM employees ORDER BY department_id;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"DISTINCT", token.DISTINCT, "DISTINCT"},
			{"column name", token.NAME, "department_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "manager_id"},
			{"comma", token.COMMA, ","},
			{"column name", token.NAME, "job_id"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"column name", token.NAME, "department_id"},
			{"semicolon", token.SEMICOLON, ";"},
		},

		// SELECT job_id AS Designation ,"Total Salary"
		// FROM employees
		// WHERE department_id<80
		// GROUP BY job_id
		// HAVING salery>10000
		// ORDER BY salery;
		{
			{"SELECT", token.SELECT, "SELECT"},
			{"column name", token.NAME, "job_id"},
			{"AS", token.AS, "AS"},
			{"column name", token.NAME, "Designation"},
			{"comma", token.COMMA, ","},
			{"string", token.STRING, "Total Salary"},
			{"FROM", token.FROM, "FROM"},
			{"table name", token.NAME, "employees"},
			{"WHERE", token.WHERE, "WHERE"},
			{"column name", token.NAME, "department_id"},
			{"lessthan", token.LESSTHAN, "<"},
			{"NUMBER", token.NUMBER, "80"},
			{"GROUP", token.GROUP, "GROUP"},
			{"BY", token.BY, "BY"},
			{"job_id", token.NAME, "job_id"},
			{"HAVING", token.HAVING, "HAVING"},
			{"salery", token.NAME, "salery"},
			{"greater than", token.GREATERTHAN, ">"},
			{"NUMBER", token.NUMBER, "10000"},
			{"ORDER", token.ORDER, "ORDER"},
			{"BY", token.BY, "BY"},
			{"salery", token.NAME, "salery"},
			{"semicolon", token.SEMICOLON, ";"},
		},
	}

	for indexInput, input := range inputs {
		test := tests[indexInput]
		lex := New(input)

		for j, tt := range test {
			tk := lex.NextToken()

			if tk.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d][%d] incorrect literal. expected=%s, got=%s", indexInput, j, tt.expectedLiteral, tk.Literal)
			}

			if tk.Type != tt.expectedType {
				t.Fatalf("tests[%d][%d] incorrect type. expected=%s, got=%s", indexInput, j, tt.expectedType, tk.Type)
			}
		}
	}
}
