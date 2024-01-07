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
								"PRIMARY",
								"KEY",
								"NOT",
								"NULL",
							},
						},
					},
				},
			},
			"CREATE TABLE orders( price INT PRIMARY KEY NOT NULL )",
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("tests[%d]", index), func(t *testing.T) {
			testString(t, test.program, test.expectedString, index)
		})
	}
}
