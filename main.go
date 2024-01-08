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
	errorJsons, _ := json.Marshal(errs)
	fmt.Print(string(errorJsons))
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
		str, _ := json.Marshal("error in parsing program")
		fmt.Print(string(str))
		os.Exit(-1)
	} else {
		str, _ := json.Marshal("OK")
		fmt.Print(string(str))
	}
}
