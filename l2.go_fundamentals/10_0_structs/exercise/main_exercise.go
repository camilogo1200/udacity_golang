package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func (c Classroom) String() string {
	return fmt.Sprintf("{\n id : %s,\n capacity : %d,\n subject : %s,\n student List : %v\n}",
		strconv.Itoa(c.id), c.capacity, c.subject, c.studentList)
}

func main() {

	c1 := Classroom{}

	var lStudent []Student
	lStudent = append(lStudent, Student{id: 1, name: "Cristhian Gómez"})
	lStudent = append(lStudent, Student{id: 2, name: "Camilo Narváez"})
	c1.id = 1001
	c1.capacity = 12
	c1.subject = "Math"
	c1.studentList = lStudent

	fmt.Println(c1)

	//& is the pointer to the memory location
	classroom2 := new(Classroom)

	fmt.Printf("Initial vars classrooom 2 %v", classroom2)

	(*classroom2).capacity = 10
	(*classroom2).id = 1102
	(*classroom2).subject = "Computer Programming"
	(*classroom2).studentList = lStudent

	fmt.Printf("Edited vars classrooom 2 %v", classroom2)

}
