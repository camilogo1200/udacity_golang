package main

import (
	"fmt"
	"strconv"
)

// FizzBuzz
func main() {
	for i := 1; i <= 15; i++ {
		var output string = ""
		if i%3 == 0 {
			output += string("Fizz")
		}
		if i%5 == 0 {
			output += string("Buzz")
		}
		if output == "" {
			fmt.Printf("%d ", i)
		} else {
			fmt.Print(output + " ")
		}
	}

	var myArr [10]string
	for i := 0; i < 10; i++ {
		myArr[i] = string(" my arr st " + strconv.Itoa(i))
		fmt.Println(myArr[i])
	}

	for i, str := range myArr {
		fmt.Println("index => " + strconv.Itoa(i) + str)
	}

}
