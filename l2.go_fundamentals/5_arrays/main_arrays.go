package main

import (
	"fmt"
)

func main() {

	// declaring array
	var vegetables [2]string
	cars := [4]string{"Mazda", "Mercedes", "BMW", "Ford"}

	vegetables[0] = "Onion"
	vegetables[1] = "Carrot"

	fmt.Println(vegetables)
	fmt.Println(cars)

	p := cars

	fmt.Println(p)
	fmt.Println(cars[2:])

	for i := 0; i < len(cars); i++ {
		fmt.Printf("cars[%d]: %s\n", i, cars[i])
	}

}
