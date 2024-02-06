package main

import (
	"QV/lexer"
	"QV/parser"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rs/cors"
)

var (
	OK = true
)

type Query struct {
	Query string `json:"query"`
}

type Jsonify interface {
	json()
}

type ErrorResponse struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Errors  []string `json:"errors"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (errorResponse ErrorResponse) json() {}

func (sucessResponse SuccessResponse) json() {}

func executeQV(query *Query) Jsonify {
	lex := lexer.New(query.Query)
	pars := parser.New(lex)
	program := pars.Parse()

	if program == nil {
		OK = false
		return ErrorResponse{
			Message: "can not parse this query",
			Status:  http.StatusBadRequest,
			Errors:  pars.Errors(),
		}
	}

	if len(pars.Errors()) != 0 {
		OK = false
		return ErrorResponse{
			Message: "invalid query",
			Status:  http.StatusBadRequest,
			Errors:  pars.Errors(),
		}
	}
	OK = true
	return SuccessResponse{
		Message: "valid query",
		Status:  http.StatusOK,
	}
}

func checkQueryHandler(writer http.ResponseWriter, request *http.Request) {
	var query Query
	err := json.NewDecoder(request.Body).Decode(&query)
	if err != nil {
		writeJSONResponse(writer, http.StatusInternalServerError, ErrorResponse{
			Message: "internal server error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	fmt.Println(query.Query)

	parserResponse := executeQV(&query)
	if !OK {
		writeJSONResponse(writer, http.StatusBadRequest, parserResponse)
		return
	}

	writeJSONResponse(writer, http.StatusOK, parserResponse)
}

func writeJSONResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		// Handle JSON encoding error
		fmt.Println(fmt.Errorf("error encoding JSON response: %v", err))
	}
}

func main() {
	mux := pat.New()

	// Set up the route for /check/ with POST method
	mux.Post("/check", http.HandlerFunc(checkQueryHandler))

	mux.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writeJSONResponse(writer, http.StatusNotFound, ErrorResponse{
			Message: "route not found",
			Status:  http.StatusNotFound,
		})
	})

	// Use the cors.Default() method to create a new CORS handler
	corsHandler := cors.Default().Handler(mux)

	// Start the server with the CORS handler
	http.ListenAndServe(":8080", corsHandler)
}
