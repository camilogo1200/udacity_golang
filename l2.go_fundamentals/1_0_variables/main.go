package main

import "fmt"

func main() {
	var product string = "T - shirt"
	var cost = 20
	fmt.Println("Product's value is:", product)
	fmt.Printf("product's type: %T\n", product)
	fmt.Println("product's cost:", cost)
	fmt.Printf("cost type: %T\n", cost)

	product2 := "T-shirt"
	cost2 := 15
	fmt.Println("Product's value is:", product2)
	fmt.Printf("product's type: %T\n", product2)
	fmt.Printf("product's type: %v\n", product2)
	fmt.Printf("product's type: %#v\n", product2)
	fmt.Println("product's cost:", cost2)
	fmt.Printf("cost type: %T\n", cost2)

	var booleanVar = true
	fmt.Printf("Boolean value = %t\n", booleanVar)
	fmt.Printf("Boolean value = %v\n", booleanVar)
	fmt.Printf("Boolean value = %#v\n", booleanVar)
	fmt.Println("Boolean value =", booleanVar)

}
