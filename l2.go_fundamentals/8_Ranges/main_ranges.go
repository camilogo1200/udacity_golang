package main

import (
	"fmt"
	"math/rand"
)

func main() {
	athletes := []string{"Camilo", "Claire", "Stephen", "Harrison", "Andrew"}
	for index, name := range athletes {
		fmt.Printf("index = %d, name = %s\n", index, name)
	}

	numbersSlice := generateSliceNumbers()
	printSlice(numbersSlice)
	sumUpSlice := sum(numbersSlice)
	fmt.Printf("total sum of elements on slice is => %d", sumUpSlice)
}

func sum(numbersSlice []int) int {
	var totalSum = 0
	for _, number := range numbersSlice {
		totalSum += number
	}
	return totalSum
}

func printSlice(slice []int) {
	for index, number := range slice {
		fmt.Printf("(%d - [%d])\n", index, number)
	}
	fmt.Println()
}

func generateSliceNumbers() []int {
	sliceLength := rand.Intn(15)
	var newSlice []int
	for i := 0; i < sliceLength; i++ {
		newSlice = append(newSlice, rand.Intn(10))
	}
	return newSlice
}
