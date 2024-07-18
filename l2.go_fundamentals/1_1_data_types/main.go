package main

func main() {
	/*
	 Numeric types
	*/
	//unsigned integers
	var uIntVar1 uint8  //0 - 255
	var uIntVar2 uint16 //0 - 65535
	var uIntVar3 uint32 //0 - 4294967295
	var uIntVar4 uint64 //0 - 18446744073709551615

	//integers
	var intVar1 int8  // -128 - 127
	var intVar2 int16 //-32768 - 32767
	var intVar3 int32 // -2147483648 - 2147483647
	var intVar4 int64 //-9223372036854775808 - 9223372036854775807

	//precision numbers
	var flt1 float32 //32-bit IEEE-754
	var flt2 float64 //64-bit IEEE-754

	var cplx3 complex64
	var cplx4 complex128

	var bty byte // byte is the alias for uint8
	var rn rune  //alias for int32

	//set of predeclared integer types
	var ui uint //either 32 or 64 bits depending on the impl or the platform
	var i int   // same size as uint

	var uPtr uintptr // an unsigned integer large enough to store the uninterpreted bits of a pointer value

	/*
		String types
					-> possibly empty  sequence of bytes
					the number of bytes is the length of the string and is never negative
		Strings are Immutable:
					-> Once created, it is impossible to change the contents of a string
					The length of a string can be discovered using the "len" function
					string's bytes can be accessed by integer indices 0 through len(s) -1.
					it is illegal to take the address of the element s[i], &s[i] is invalid
	*/

	var name string = "John Doe"
	nLength := len(name)

	/*
			Array Types
				-> An array is a numbered sequence of elements of a single type, called the element type
		 			The number of elements is called the length of the array and is never negative

			ArrayType = [ arrayLength ] ElementType
				arrayLength = Expression
				elementType = Type
	*/

	var arr [32]byte
	var arr2 [32]*float64 //array of pointers
	var matrix [3][5]int
	var matrix1 [3][3][3]int
	var funcArray [10]func() int

}
