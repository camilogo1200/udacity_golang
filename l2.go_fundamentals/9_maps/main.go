package main

import "fmt"

func main() {

	var myNewMap map[string]string
	myNewMap = make(map[string]string)

	//initialize with this shortcut
	myNewMap2 := map[string]string{}

	myNewMap["Hello"] = "World!"
	hello := myNewMap["Hello"]
	fmt.Printf("Hello => %s\n", hello)
	world := myNewMap["World"]
	fmt.Printf("World => %s\n", world)
	fmt.Printf("World => %t\n", world == "")

	myNewMap["a"] = "hello1"
	myNewMap["b"] = "hello2"
	myNewMap["c"] = "hello3"
	myNewMap["d"] = "hello4"
	myNewMap["e"] = "hello5"

	fmt.Printf("map total length => %d\n", len(myNewMap))
	fmt.Println(myNewMap)
	fmt.Println(myNewMap2)

	fmt.Println("Removed elements on map")
	delete(myNewMap, "Hello")
	delete(myNewMap, "World")

	fmt.Printf("map total length => %d\n", len(myNewMap))
	fmt.Println(myNewMap)

	//test key existency

	val, exist := myNewMap["World"]
	val2, exist2 := myNewMap["a"]

	fmt.Printf("Existe key \"a\" => %t, value => %s\n", exist2, val2)
	fmt.Printf("Existe key \"World\" => %t value => %s\n", exist, val)

	//iterate over a map

	for key, value := range myNewMap {
		fmt.Printf("[key => %s value => %s]\n", key, value)
	}

	// Initialize a map
	commits := map[string]int{
		"abc":  1234,
		"abc2": 1234,
		"abc3": 1234,
		"abc4": 12345,
	}
	for key, value := range commits {
		fmt.Printf("[key => %s value => %d]\n", key, value)
	}

	//add duplicated key
	commits["add"] = 1
	commits["add"] = 2
	commits["add"] = 3

	for key, value := range commits {
		fmt.Printf("[key => %s value => %d]\n", key, value)
	}

}
