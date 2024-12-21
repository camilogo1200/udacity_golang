package main

import (
	"fmt"
	"variableScope/myPackage"
)

func main() {

	blockLevelVariable := "this is a block level variable"

	fmt.Println(blockLevelVariable)

	exportedVariable := myPackage.PublicVar

	fmt.Println(exportedVariable)

	myPackage.Exported()
}
