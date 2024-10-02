package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	fmt.Println(" -> filesPath is ", filesPath)

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

		if strings.Contains(lineText, "## ") && !strings.Contains(lineText, "### ") {

			lineText = strings.TrimSpace(lineText)
			lineText = strings.ToLower(lineText)
			lineText = strings.Replace(lineText, "## ", "", -1)

			linkName := "- [" + lineText + "]"

			linkAdress := strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"

			linkForIndex := linkName + linkAdress

			outResult = fmt.Sprintf("%s%s\n", outResult, linkForIndex)
		}
	}
	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}
	fmt.Print(outResult)
}
