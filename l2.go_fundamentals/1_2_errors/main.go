package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserId : %d", a, b, os.Geteuid())
	}
	return nil
}

func main() {

	err := check(0, 2)
	if err == nil {
		fmt.Println("check executed successfully")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)

	if err != nil && err.Error() == "This is a custom error message." {
		fmt.Println("Custom error detected.")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("int value is ", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
