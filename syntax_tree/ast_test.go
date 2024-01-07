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
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("tests[%d]", index), func(t *testing.T) {
			testString(t, test.program, test.expectedString, index)
		})
	}
}
