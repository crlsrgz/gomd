package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Function start")
	// open file
	f, err := os.Open("doc.md")

	if err != nil {
		log.Fatal(err)
	}

	//close the file at the end
	defer f.Close()

	// read line by line
	scanner := bufio.NewScanner(f)

	outResult := ""

	var userInput string
	fmt.Println("Give me input: ")
	fmt.Scan(&userInput)

	for scanner.Scan() {
		// fmt.Printf("line Nr. : %s\n", scanner.Text())
		lineText := scanner.Text()
		if strings.Contains(lineText, "## ") {

			outResult = fmt.Sprintf("%s%s\n", outResult, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}
	fmt.Print(outResult)
	fmt.Printf("user input is: %s\n", userInput)
}
