package main

import (
	"fmt"
	"time"
)

func iterates(str []string) {
	for i, value := range str {
		fmt.Printf("[Index %d - value=%s\n]", i, value)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	slc1 := []string{"hi", "there", "!"}
	slc2 := []string{"hi", "there", "!", "you", "'re", "always", "welcome"}
	iterates(slc1)
	iterates(slc2)

	waitChannel := make(chan bool)
	go iterates(slc1)
	go iterates(slc2)

	go func() {
		time.Sleep(10 * time.Second)
		waitChannel <- true
	}()

	retWait := <-waitChannel
	fmt.Printf("RetWait : %v\n", retWait)
}
