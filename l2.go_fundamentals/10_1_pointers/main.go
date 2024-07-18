package main

import (
	"fmt"
	"image"
)

/*
	A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer.
	The value of an uninitialized pointer is "nil"

	PointerType = "*" BaseType
	BaseType = Type


	Pointers are used to update memory
*/

type TestPtrStruct struct {
	email string
	tags  []string
}

func (tPtr TestPtrStruct) Email() string {
	return tPtr.email
}

func (tPtr TestPtrStruct) Tags() []string {
	return tPtr.tags
}

func (tPtr TestPtrStruct) UpdateEmail(email string) {
	tPtr.email = email // returns a copy of the value
}

func (tPtr *TestPtrStruct) UpdateEmailPtr(email string) {
	tPtr.email = email
}

func (tPtr TestPtrStruct) addTag(tag string) {
	var address = &tPtr

	fmt.Printf("( copy Value ) struct address = %v\n", address)

	if tPtr.tags == nil {
		tPtr.tags = make([]string, 10)
	}

	tPtr.tags = append(tPtr.tags, tag)
}

func (tPtr *TestPtrStruct) addTagPtr(tag string) {
	var address = &tPtr
	fmt.Printf("( ref Value ) struct address = %v\n", address)
	if tPtr.tags == nil {
		tPtr.tags = make([]string, 0)
	}

	tPtr.tags = append(tPtr.tags, tag)
}

func main() {
	var myPtr *image.Point = &image.Point{
		X: 0,
		Y: 0,
	}

	arr := [4]string{"a", "b", "c", "d"}
	var myPtr1 *[4]string = &arr
	fmt.Println(myPtr, myPtr1)
	testPtr := TestPtrStruct{email: "email@myemail.com", tags: nil}
	fmt.Println(testPtr)
	//emails does not change... the struct passed to the function is by value
	//we will need to use the pointer to update the memory of the struct on email field
	testPtr.UpdateEmail("MyNewEmail@newEmail.com")
	fmt.Println(testPtr)
	testPtr.UpdateEmailPtr("MyNewEmail@mewEmail.com")
	fmt.Println(testPtr)
	//does not work we need to use Pointers to add / update the memory
	testPtr.addTag("newTag")
	fmt.Println(testPtr)
	testPtr.addTagPtr("ptrTag")
	fmt.Println(testPtr)
}
