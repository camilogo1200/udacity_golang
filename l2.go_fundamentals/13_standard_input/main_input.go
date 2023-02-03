package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//readSingleCharacter()
	//readSingleLine()
	//readMultipleLines()
	//readFormattedUserInput()
	//readNumbers()
}

func readSingleCharacter() {
	fmt.Println("Input text:")
	/*
		//bufio.reader
		reader := bufio.NewReader(os.Stdin)
		char, size, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
	*/

	var char rune
	n, err := fmt.Scanf("%c", &char)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read character => %c, read elements => %d-\n", char, n)
}

// readSingleLine is a function for reading only a single line of text from the os.Stdin
func readSingleLine() {
	var name string

	fmt.Println("insert your name")
	//Use fmt.Scan when you want to read each word of a line into a different variable
	// items, error := fmt.Scan(&name, &lastname)
	reader := bufio.NewReader(os.Stdin)
	//Use bufio.Reader if you want to read a full line including the newline character
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name  => %s ", name)
	fmt.Printf("error => %v", err)

	//USe bufio.Scanner when you want to read a full line of text without a newline character
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	errorScann := scanner.Err()
	if errorScann != nil {
		log.Fatal(errorScann)
	}

	readStr := scanner.Text()
	fmt.Printf("Read Line: %s\n", readStr)
}

func readMultipleLines() {

	fmt.Println("input text:")
	/*
		//Fmt.Scan
		var w1, w2, w3 string
		n, err := fmt.Scan(&w1, &w2, &w3)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("number of items read: %d\n", n)
		fmt.Printf("read text: %s %s %s-\n", w1, w2, w3)
	*/

	/*
		//bufio.Reader multiline

		reader := bufio.NewReader(os.Stdin)

		var lines []string

		for {
			line, errReader := reader.ReadString('\n')
			if errReader != nil {
				log.Fatal(errReader)
			}

			if len(strings.TrimSpace(line)) == 0 {
				break
			}
			lines = append(lines, line)
			fmt.Printf("line read : %s", line)
		}

		fmt.Printf("total lines => %d\n", len(lines))
		fmt.Println("read lines-> \n", lines)
	*/
	//bufio.scanner

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			log.Fatal(err)

		}
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Printf("total lines => %d\n", len(lines))
	for index, line := range lines {
		fmt.Printf("index -> %d, line -> %s\n", index, line)
	}
}

func readFormattedUserInput() {
	fmt.Println("input text:")
	var name string
	var country string
	//must be exact match
	n, err := fmt.Scanf("%s is born in %s", &name, &country)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Println(name, country)
}

func readNumbers() {
	fmt.Println("Input Number:")
	var number int64
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read number: %d\n", number)

	num2 := number + 1001
	fmt.Printf("new number: %d\n", num2)
}
