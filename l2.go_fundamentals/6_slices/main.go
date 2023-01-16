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
	fmt.Printf("Slice length => %d,  capacity => %d\n", len(animals), cap(animals))

	//declare empty slices
	var newSlice []int
	fmt.Printf(" initial Slice length => %d,  capacity => %d\n", len(newSlice), cap(newSlice))

	newSlice = append(newSlice, 1)
	newSlice = append(newSlice, 2)
	newSlice = append(newSlice, 3)
	newSlice = append(newSlice, 4)

	fmt.Printf(" Slice length => %d,  capacity => %d\n", len(newSlice), cap(newSlice))

	fmt.Printf("new slice => %d\n", newSlice)
	fmt.Printf("new slice => %v\n", newSlice)

	var myNewSlice2 []int
	myNewSlice2 = make([]int, 10)

	fmt.Printf(" initial Slice length => %d,  capacity => %d\n", len(myNewSlice2), cap(myNewSlice2))

}
