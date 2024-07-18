package main

import (
	"fmt"
	"runtime"
)

func main() {
	language := "Go"

	if language == "Java" {
		fmt.Printf("Language is \"Java\"")
	} else if language == "Rust" {
		fmt.Println("Language is \"Rust\"")
	} else if language == "C++" {
		fmt.Println("Language is \"C++\"")
	} else {
		fmt.Println("Language is not Java, Rust or C++")
	}

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
