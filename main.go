package main

import (
	"bufio"
	"fmt"
	"gocli/readlines"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Function start")

	fmt.Println("give me the file path: ")
	filesPath := ""
	fmt.Scanln(&filesPath)

	// open file
	f, err := os.Open(filesPath)

	if err != nil {
		log.Fatal(err)
	}

	//close the file at the end
	defer f.Close()

	// read line by line
	scanner := bufio.NewScanner(f)

	outResult := ""

	for scanner.Scan() {
		// fmt.Printf("line Nr. : %s\n", scanner.Text())
		lineText := scanner.Text()
		if strings.Contains(lineText, "## ") && !strings.Contains(lineText, "### ") {

			outResult = fmt.Sprintf("%s%s\n", outResult, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}
	fmt.Print(outResult)
	readlines.Readlines()
}
