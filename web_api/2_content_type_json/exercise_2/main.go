package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var cityPopulations = map[string]float64{
	"Tokio":       13960000,
	"Delhi":       32941000,
	"Shanghai":    26320000,
	"Sao Paulo":   12330000,
	"Mexico City": 8855000}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Starting server on port : 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cityPopulations)
}
