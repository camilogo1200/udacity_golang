package main

import "fmt"

func main() {
	var language = "Go"
	const released = 2009
	isProgrammingLanguage := true

	fmt.Printf("language = %s\n", language)
	fmt.Printf("language = %q\n", language)
	fmt.Printf("language = %x\n", language)
	fmt.Printf("language = %X\n", language)
	fmt.Printf("release = %d\n", released)
	fmt.Printf("isProgrammingLanguage = %t\n", isProgrammingLanguage)
	fmt.Println(language, released, isProgrammingLanguage)

}
