package main

import (
	"QV/lexer"
	"QV/parser"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	ERROR = "ERROR"
	OK    = "OK"
)

type Query struct {
	input string
}

func showParserErrors(parser *parser.Parser) {
	errs := parser.Errors()
	if len(errs) == 0 {
		return
	}
	fmt.Printf("%s:", ERROR)
	errorJsons, _ := json.Marshal(errs)
	fmt.Println(string(errorJsons))
	os.Exit(-1)
}

func main() {
	var query Query

	// Parsing input query and store it in Query.input
	flag.StringVar(&query.input, "q", "CREATE TABLE test (id INT PRIMARY KEY);", "CREATE TABLE Query that want to validate")
	flag.Parse()

	lex := lexer.New(query.input)
	pars := parser.New(lex)
	program := pars.Parse()
	showParserErrors(pars)

	if program == nil {
		fmt.Println(fmt.Sprintf("%s:%s", ERROR, "error in parsing program"))
		os.Exit(-1)
	} else {
		fmt.Println(fmt.Sprintf("%s:%s", OK, program.String()))
	}
}
