package main

import "fmt"

/*
 	Struct types
	A Struct is a sequence of named elements, called fields, each of which has a name and a type.
	Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
	Within a struct, non-blank field must be unique.

	StructType = "struct {" { FieldDeclaration ";" } "}"
	FieldDeclaration = (IdentifierList type | EmbeddedField) [Tag]
	EmbeddedField = ["*"] TypeName [ TypeArgs ]
	Tag = string_literal

*/

type myStruct struct {
	x, y float64 "" // empty tag string is like an absent tag
	name string  "any string is  permitted as a tag"
}

//A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers
//they follow the convention outlined byt the reflect package

type MyProtobufStruct struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}

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
