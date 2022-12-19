package main

import "fmt"

func main() {
	athletes := []string{"Camilo", "Claire", "Stephen", "Harrison", "Andrew"}

	for index, name := range athletes {
		fmt.Printf("index = %d, name = %s\n", index, name)
	}

}
