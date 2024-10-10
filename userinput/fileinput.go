package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Fileinput() {

	fmt.Println("-> Give me the file path: ")
	filesPath := ""
	_, errPath := fmt.Scanln(&filesPath)
	if errPath != nil {
		log.Fatal(errPath)
	}

	filesPath = filepath.ToSlash(filesPath)

	if filesPath[:1] == "/" {
		filesPath = filesPath[1:]
	}

	fmt.Printf("--> filesPath is %s\n", filesPath)

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
		lineText := scanner.Text()
		if lineText != "" && lineText != "\n" {
			outResult = outResult + BuildIndexList(lineText)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(outResult)
}
