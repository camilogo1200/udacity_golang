package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There! Current time: %v", time.Now().Format(time.DateTime))
}

const LISTEN_PORT = 3500

func main() {
	http.HandleFunc("/", index)
	fmt.Println("starting server...")
	fmt.printf("Listening on Port => ")
	http.ListenAndServe(":", nil)
}
