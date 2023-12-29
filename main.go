package main

import (
	"flag"
)

type Query struct {
	input string
}

func main() {
	var query Query

	flag.StringVar(&query.input, "q", "CREATE TABLE test (id INT);", "CREATE TABLE Query that want to validate")
	flag.Parse()
}
