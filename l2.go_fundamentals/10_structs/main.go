package main

import "fmt"

type Car struct {
	make string
	year uint16
	used bool
}

func (c Car) getMaker() string {
	return c.make
}

func main() {
	car1 := Car{make: "Nissan Murano", year: 2019, used: true}
	car2 := Car{make: "Audi R8", year: 2023, used: false}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.getMaker())


}
