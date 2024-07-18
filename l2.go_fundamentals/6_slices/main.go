package main

import "fmt"

/*
	Slice types
		-> A slice is a descriptor for a contiguous segment of an underlying array and provides access
			to a numbered sequence of elements from that array.
			A slice type denotes the set of all slices of arrays of its element type.
		-> The number of elements is called the length og the slice and is never negative.
		-> The value of an uninitialized slice is "nil"

		SliceType = [] ElementType

		-> The length of a slice can be discovered by the built-in function "len", this can be changed during execution
		-> The elements can be addressed by integer indices 0 through len(s) -1.
		-> The slice index may be less than the index of the same element in the underlying array.
		-> A slice once initialized, is always associated with an underlying array, that holds its elements.
			A slice therefore,  shares storage with its array and with other slices of the same array.
			by contrast distinct arrays always represent distinct storage
		-> The array underlying the slice may extend past the end of the slice.
		-> The capacity is a measure of that extent: is the sum of the length of the slice and the length of the array
			beyond the slice.
		-> The capacity of a slice can be discovered using the built-in function "cap"
		-> A new, initialized slice value for a given element type T may be made using the built-in function "make"
			which takes a slice type, and parameters specifying the length and optionally the capacity.
		-> A slice allocated with "make" always allocates a new, hidden array to which the returned slice value refers.
		-> executing :
			make ( []T , length, capacity )

		-> slices are always one-dimensional but may be composed to create a higher-dimensional objects.
			** with slice of slices ( or array of slices ), the inner lengths may vary dynamically,
				Moreover, the inner Slices MUST BE initialized individually
*/

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
