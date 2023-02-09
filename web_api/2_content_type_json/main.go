package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var dictionary = map[string]string{}

func main() {

	http.HandleFunc("/", returnJsonDictionary)
	http.HandleFunc("/not-found", handleNotFoundStatus)
	fmt.Println("Starting servenr on port : 3000")
	http.ListenAndServe(":3000", nil)
}

func returnJsonDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getDictionary())
}

func handleNotFoundStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func getDictionary() map[string]string {
	dictionary["golang"] = "Golang language"
	dictionary["creation_date"] = "2009"
	dictionary["version"] = "1.20"
	dictionary["release_Date"] = "02/01/2023"
	return dictionary
}
