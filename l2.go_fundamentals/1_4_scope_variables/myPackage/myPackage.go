package myPackage

import "fmt"

var privateVar = "THis is a private var and is not exported"

var PublicVar = "Public variable (Exported)"

func notExported() {
	fmt.Println(privateVar)
}

func Exported() {
	fmt.Printf("From my package : %s ", PublicVar)
}
