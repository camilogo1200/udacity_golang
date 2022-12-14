package main

import "fmt"

func main() {

	animals := []string{"Bear", "Dog", "Cat", "Bird"}

	fmt.Println(animals)
	fmt.Printf("Animals slice len => %d", len(animals))
	fmt.Println(animals[1:])
	animals2 := animals
	animals = append(animals, "Polar Bear", "Fox", "Polar Fox")
	fmt.Printf("Animals slice len => %d", len(animals))
	fmt.Println(animals)
	fmt.Printf("Animals2 slice len => %d", len(animals2))
	fmt.Println(animals2)
}
