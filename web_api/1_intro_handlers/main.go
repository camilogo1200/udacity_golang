package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There! Current time: %v", time.Now().Format(time.DateTime))
}

const LISTEN_PORT = 3500

func main() {
	http.HandleFunc("/", index)
	fmt.Println("starting server...")
	fmt.Printf("Listening on Port => %v\n", LISTEN_PORT)
	serverString := ":" + strconv.Itoa(LISTEN_PORT)
	http.ListenAndServe(serverString, nil)

}
