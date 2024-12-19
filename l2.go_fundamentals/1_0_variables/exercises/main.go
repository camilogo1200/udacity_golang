package main

import "fmt"

func main() {
	//first way of declaration for a variable
	var myString string
	myString = "Hello World"
	fmt.Println(myString)

	//second way to declare a variable
	var myNumber int = 33
	fmt.Println(myNumber)

	//Third way of declaring a variable
	mySlice := []int{1, 2, 3}
	println(mySlice[0])

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
