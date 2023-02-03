package main

import "fmt"

type Football struct{}

type Baseball struct {
	mass         int
	acceleration int
}
type Force interface {
	getForce() int
}

func (b Baseball) getForce() int {
	return b.mass * b.acceleration
}

func (f Football) getForce() int {
	return 50000
}

func compareForces(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func main() {
	var b1 Baseball
	b1 = Baseball{mass: 2, acceleration: 20}
	f1 := Football{}

	fmt.Println(compareForces(f1, b1))
}
