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
		CREATE TABLE test12 ( col VARCHAR );
		`,
		`
		CREATE TABLE hello ( cl1 INT );
		`,
		`
		CREATE TABLE _name_start_with_under_line (
			is_ok BOOLEAN,
			col2 VARCHAR
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
		CREATE TABLE IF NOT EXISTS test_test_test ( col FLOAT );
		`,
		`
		CREATE TABLE Customers (
			First_Name VARCHAR NOT NULL,
			Last_Name VARCHAR NOT NULL,
			City VARCHAR NOT NULL,
			Email VARCHAR NOT NULL,
			Phone_Number VARCHAR NOT NULL,
			Registration_Date DATE NOT NULL
			);
		`,

		`
		CREATE TABLE Customers (
			First_Name VARCHAR NOT NULL,
			Last_Name VARCHAR NOT NULL,
			Email VARCHAR NOT NULL,
			Phone_Number VARCHAR NOT NULL
			);
		`,

		`
		CREATE TABLE books (
			book_id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR NOT NULL,
			author VARCHAR NOT NULL,
			genre VARCHAR NOT NULL,
			publication_date DATE NOT NULL
			);
		`,

		`
		CREATE TABLE products (
			product_id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR NOT NULL,
			price DECIMAL NOT NULL,
			description TEXT,
			stock_quantity INT NOT NULL
			);
		`,

		`
		CREATE TABLE employees (
			employee_id INT PRIMARY KEY AUTO_INCREMENT,
			first_name VARCHAR NOT NULL,
			last_name VARCHAR NOT NULL,
			email VARCHAR UNIQUE NOT NULL,
			phone_number VARCHAR NOT NULL,
			department VARCHAR NOT NULL
			);
		`,

		`
		CREATE TABLE customers (
			customer_id INT PRIMARY KEY AUTO_INCREMENT,
			first_name VARCHAR NOT NULL,
			last_name VARCHAR NOT NULL,
			email VARCHAR UNIQUE NOT NULL,
			address VARCHAR NOT NULL,
			city VARCHAR NOT NULL,
			state VARCHAR NOT NULL,
			zip_code VARCHAR NOT NULL
			);
		`,

		`
		CREATE TABLE orders (
			order_id INT PRIMARY KEY AUTO_INCREMENT,
			customer_id INT NOT NULL,
			order_date DATE NOT NULL,
			total_amount DECIMAL NOT NULL,
			shipping_address VARCHAR,
			billing_address VARCHAR
			);
		`,

		`
		CREATE TABLE order_items (
			order_item_id INT PRIMARY KEY AUTO_INCREMENT,
			order_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL,
			unit_price DECIMAL NOT NULL
			);
		`,

		`
		CREATE TABLE users (
			user_id INT PRIMARY KEY AUTO_INCREMENT,
			username VARCHAR UNIQUE NOT NULL,
			password VARCHAR NOT NULL,
			email VARCHAR UNIQUE NOT NULL,
			first_name VARCHAR,
			last_name VARCHAR
			);
		`,

		`
		CREATE TABLE posts (
			post_id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR NOT NULL,
			content TEXT NOT NULL,
			author_id INT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			);
		`,

		`
		CREATE TABLE comments (
			comment_id INT PRIMARY KEY AUTO_INCREMENT,
			post_id INT NOT NULL,
			user_id INT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`,

		`
		CREATE TABLE categories (
			category_id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR NOT NULL,
			description TEXT
			);
		`,

		`
		CREATE TABLE posts_categories (
			post_category_id INT PRIMARY KEY AUTO_INCREMENT,
			post_id INT NOT NULL,
			category_id INT NOT NULL
			);
		`,

		`
		CREATE TABLE settings (
			setting_id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR NOT NULL,
			value TEXT NOT NULL
			);
		`,

		`
		CREATE TABLE logs (
			log_id INT PRIMARY KEY AUTO_INCREMENT,
			message TEXT NOT NULL,
			level VARCHAR NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`,

		`
		CREATE TABLE attachments (
			attachment_id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR NOT NULL,
			MIME_type VARCHAR NOT NULL,
			file_size INT NOT NULL,
			file_content BLOB NOT NULL
			);
		`,

		`
		CREATE TABLE IF NOT EXISTS books (
			book_id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR NOT NULL,
			author VARCHAR NOT NULL,
			genre VARCHAR NOT NULL,
			publication_date DATE NOT NULL
		  );
		`,

		`
		CREATE TABLE IF NOT EXISTS products (
			product_id INT
		   
		  PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR NOT
		   
		  NULL,
			price DECIMAL NOT
		   
		  NULL,
			description TEXT,
			stock_quantity INT NOT NULL
		  );
		`,

		`
		CREATE TABLE IF NOT EXISTS employees (
			employee_id INT
		   
		  PRIMARY
		   
		  KEY AUTO_INCREMENT,
			first_name VARCHAR NOT
		   
		  NULL,
			last_name VARCHAR NOT
		   
		  NULL,
			email VARCHAR
		   
		  UNIQUE
		   
		  NOT
		   
		  NULL,
			phone_number VARCHAR NOT NULL,
			department VARCHAR NOT NULL
		  );
		`,

		`
		CREATE
		
		TABLE IF NOT
		
		EXISTS customers (
		customer_id INT
		
		PRIMARY KEY AUTO_INCREMENT,
		first_name VARCHAR NOT
		
		NULL,
		last_name VARCHAR NOT
		
		NULL,
		email VARCHAR UNIQUE NOT
		
		NULL,
		address VARCHAR NOT
		
		NULL,
		city VARCHAR NOT
		
		NULL,
		state VARCHAR NOT
		
		NULL,
		zip_code VARCHAR NOT
		
		NULL
		);
		`,

		`
		CREATE TABLE IF NOT EXISTS order_items (
			order_item_id INT PRIMARY KEY AUTO_INCREMENT,
			order_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL,
			unit_price DECIMAL NOT NULL
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
		{"books"},
		{"products"},
		{"employees"},
		{"customers"},
		{"orders"},
		{"order_items"},
		{"users"},
		{"posts"},
		{"comments"},
		{"categories"},
		{"posts_categories"},
		{"settings"},
		{"logs"},
		{"attachments"},
		{"books"},
		{"products"},
		{"employees"},
		{"customers"},
		{"order_items"},
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
