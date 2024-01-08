package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/bmizerany/pat"
)

type Query struct {
	Query string `json:"query"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func executeQV(query *Query) (string, error) {
	// fmt.Println(`-q="` + query.Query + `"`)
	cmd := exec.Command("./../../bin/linux_amd64/QV", `-q="`+query.Query+`"`)

	output, err := cmd.Output()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
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

	str, err := executeQV(&query)
	if err != nil {
		writeJSONResponse(writer, http.StatusBadRequest, ErrorResponse{
			Message: str,
			Status:  http.StatusBadRequest,
		})
		return
	}

	if str[:6] == "ERROR" {
		writeJSONResponse(writer, http.StatusBadRequest, ErrorResponse{
			Message: str[6:],
			Status:  http.StatusBadRequest,
		})
	} else {
		writeJSONResponse(writer, http.StatusOK, SuccessResponse{
			Message: str,
			Status:  http.StatusOK,
		})
	}
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

	mux.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writeJSONResponse(writer, http.StatusNotFound, ErrorResponse{
			Message: "route not found",
			Status:  http.StatusNotFound,
		})
	})

	// Start the server
	http.ListenAndServe(":8080", mux)
}
