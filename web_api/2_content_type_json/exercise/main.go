package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	w.Header().Set("Content Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of most populous cities: <br/>")

	for city, population := range cityPopulations {
		fmt.Fprintf(w, "<h2> %s</h2>", city)
		fmt.Fprintf(w, "Population : %s <br/>", strconv.FormatFloat(population, 'f', 5, 64))
	}

}
