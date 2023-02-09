package main

import (
	"fmt"
	"net/http"
	"time"
)

var cities = []string{"Tokio", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main Page\nWelcome => %v", time.RFC1123)
}

func citylist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most Popular Cities :\n")
	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/citylist", citylist)
	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", nil)
}
