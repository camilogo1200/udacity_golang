package main

import (
	"fmt"
	"runtime"
)

func main() {
	languague := "Go"

	if languague == "Java" {
		fmt.Printf("Language is \"Java\"")
	} else if languague == "Rust" {
		fmt.Println("Language is \"Rust\"")
	} else if languague == "C++" {
		fmt.Println("Language is \"C++\"")
	} else {
		fmt.Println("Language is not Java, Rust or C++")
	}

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
