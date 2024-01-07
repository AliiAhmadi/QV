package ast

import (
	"QV/token"
	"fmt"
	"testing"
)

func testString(t *testing.T, program *Program, expectedString string, index int) {
	givenString := program.String()

	if expectedString != givenString {
		t.Fatalf("invalid query in [%d] - error in strigifed program", index)
	}
}

func TestStringifingPrograms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		program        *Program
		expectedString string
	}{
		{
			&Program{
				Query: CreateQuery{
					Token: token.Token{
						Type:    token.CREATE,
						Literal: token.CREATE,
					},
					Name: &Identifier{
						Token: token.Token{
							Type:    token.NAME,
							Literal: "orders",
						},
						Value: "orders",
					},
					Columns: []Column{
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "price",
								},
								Value: "price",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.PRIMARY,
										Literal: token.PRIMARY,
									},
								},
								{
									Token: token.Token{
										Type:    token.KEY,
										Literal: token.KEY,
									},
								},
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
					},
				},
			},
			`CREATE TABLE orders( price INT PRIMARY KEY NOT NULL )`,
		},
		{
			&Program{
				Query: CreateQuery{
					Token: token.Token{
						Type:    token.CREATE,
						Literal: token.CREATE,
					},
					Name: &Identifier{
						Token: token.Token{
							Type:    token.NAME,
							Literal: "products",
						},
						Value: "products",
					},
					Columns: []Column{
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "productID",
								},
								Value: "productID",
							},
							DataType: token.INTEGER,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.PRIMARY,
										Literal: token.PRIMARY,
									},
								},
								{
									Token: token.Token{
										Type:    token.KEY,
										Literal: token.KEY,
									},
								},
							},
						},
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "name",
								},
								Value: "name",
							},
							DataType: token.TEXT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
								{
									Token: token.Token{
										Type:    token.UNIQUE,
										Literal: token.UNIQUE,
									},
								},
							},
						},
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "description",
								},
								Value: "description",
							},
							DataType: token.TEXT,
							MetaData: []token.MetaData{},
						},
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "category",
								},
								Value: "category",
							},
							DataType: token.TEXT,
							MetaData: []token.MetaData{},
						},
					},
				},
			},
			`CREATE TABLE products( productID INTEGER PRIMARY KEY, name TEXT NOT NULL UNIQUE, description TEXT, category TEXT )`,
		},
		{
			&Program{
				Query: CreateQuery{
					Token: token.Token{
						Type:    token.CREATE,
						Literal: token.CREATE,
					},
					Name: &Identifier{
						Token: token.Token{
							Type:    token.NAME,
							Literal: "customers",
						},
						Value: "customers",
					},
					Columns: []Column{
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "customer_id",
								},
								Value: "customer_id",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.PRIMARY,
										Literal: token.PRIMARY,
									},
								},
								{
									Token: token.Token{
										Type:    token.KEY,
										Literal: token.KEY,
									},
								},
								{
									Token: token.Token{
										Type:    token.AUTO_INCREMENT,
										Literal: token.AUTO_INCREMENT,
									},
								},
							},
						},
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "first_name",
								},
								Value: "first_name",
							},
							DataType: token.VARCHAR,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
					},
				},
			},
			`CREATE TABLE customers( customer_id INT PRIMARY KEY AUTO_INCREMENT, first_name VARCHAR NOT NULL )`,
		},

		{
			&Program{
				Query: CreateQuery{
					Token: token.Token{
						Type:    token.CREATE,
						Literal: token.CREATE,
					},
					Name: &Identifier{
						Token: token.Token{
							Type:    token.NAME,
							Literal: "comments",
						},
						Value: "comments",
					},
					// comment_id INT PRIMARY KEY AUTO_INCREMENT
					Columns: []Column{
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "comment_id",
								},
								Value: "comment_id",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.PRIMARY,
										Literal: token.PRIMARY,
									},
								},
								{
									Token: token.Token{
										Type:    token.KEY,
										Literal: token.KEY,
									},
								},
								{
									Token: token.Token{
										Type:    token.AUTO_INCREMENT,
										Literal: token.AUTO_INCREMENT,
									},
								},
							},
						},
						// post_id INT NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "post_id",
								},
								Value: "post_id",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// user_id INT NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "user_id",
								},
								Value: "user_id",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// content TEXT NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "content",
								},
								Value: "content",
							},
							DataType: token.TEXT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "created_at",
								},
								Value: "created_at",
							},
							DataType: token.DATETIME,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
								{
									Token: token.Token{
										Type:    token.DEFAULT,
										Literal: token.DEFAULT,
									},
								},
								{
									Token: token.Token{
										Type:    token.CURRENT_TIMESTAMP,
										Literal: token.CURRENT_TIMESTAMP,
									},
								},
							},
						},
					},
				},
			},
			`CREATE TABLE comments( comment_id INT PRIMARY KEY AUTO_INCREMENT, post_id INT NOT NULL, user_id INT NOT NULL, content TEXT NOT NULL, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP )`,
		},

		{
			&Program{
				Query: CreateQuery{
					Token: token.Token{
						Type:    token.CREATE,
						Literal: token.CREATE,
					},
					Name: &Identifier{
						Token: token.Token{
							Type:    token.NAME,
							Literal: "books",
						},
						Value: "books",
					},
					// book_id INT PRIMARY KEY AUTO_INCREMENT
					Columns: []Column{
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "book_id",
								},
								Value: "book_id",
							},
							DataType: token.INT,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.PRIMARY,
										Literal: token.PRIMARY,
									},
								},
								{
									Token: token.Token{
										Type:    token.KEY,
										Literal: token.KEY,
									},
								},
								{
									Token: token.Token{
										Type:    token.AUTO_INCREMENT,
										Literal: token.AUTO_INCREMENT,
									},
								},
							},
						},
						// title VARCHAR NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "title",
								},
								Value: "title",
							},
							DataType: token.VARCHAR,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// author VARCHAR NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "author",
								},
								Value: "author",
							},
							DataType: token.VARCHAR,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// genre VARCHAR NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "genre",
								},
								Value: "genre",
							},
							DataType: token.VARCHAR,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
							},
						},
						// publication_date DATE NOT NULL
						{
							Token: token.Token{
								Type:    token.COLUMN,
								Literal: token.COLUMN,
							},
							Name: &Identifier{
								Token: token.Token{
									Type:    token.NAME,
									Literal: "publication_date",
								},
								Value: "publication_date",
							},
							DataType: token.DATE,
							MetaData: []token.MetaData{
								{
									Token: token.Token{
										Type:    token.NOT,
										Literal: token.NOT,
									},
								},
								{
									Token: token.Token{
										Type:    token.NULL,
										Literal: token.NULL,
									},
								},
								{
									Token: token.Token{
										Type:    token.DEFAULT,
										Literal: token.DEFAULT,
									},
								},
								{
									Token: token.Token{
										Type:    token.CURRENT_TIMESTAMP,
										Literal: token.CURRENT_TIMESTAMP,
									},
								},
							},
						},
					},
				},
			},
			`CREATE TABLE books( book_id INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR NOT NULL, author VARCHAR NOT NULL, genre VARCHAR NOT NULL, publication_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP )`,
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("tests[%d]", index), func(t *testing.T) {
			testString(t, test.program, test.expectedString, index)
		})
	}
}
